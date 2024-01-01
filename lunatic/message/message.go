// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package message provides the Go bindings to the lunatic::message API.
// From: https://github.com/lunatic-solutions/lunatic/blob/40b4b7e4e9b76e4a0b74ad2d166660dc560214b6/crates/lunatic-messaging-api/src/lib.rs#L50-L116
//
// There are two kinds of messages a lunatic process can receive:
//
//  1. **Data message** that contains a buffer of raw `u8` data and host side resources.
//  2. **LinkDied message**, representing a `LinkDied` signal that was turned into a message. The
//     process can control if when a link dies the process should die too, or just receive a
//     `LinkDied` message notifying it about the link's death.
//
// All messages have a `tag` allowing for selective receives. If there are already messages in the
// receiving queue, they will be first searched for a specific tag and the first match returned.
// Tags are just `i64` values, and a value of 0 indicates no-tag, meaning that it matches all
// messages.
//
// # Data messages
//
// Data messages can be created from inside a process and sent to others.
//
// They consists of two parts:
// * A buffer of raw data
// * An collection of resources
//
// If resources are sent between processes, their ID changes. The resource ID can for example
// be already taken in the receiving process. So we need a way to communicate the new ID on the
// receiving end.
//
// When the `create_data(tag, capacity)` function is called an empty message is allocated and both
// parts (buffer and resources) can be modified before it's sent to another process. If a new
// resource is added to the message, the index inside of the message is returned. This information
// can be now serialized inside the raw data buffer in some way.
//
// E.g. Serializing a structure like this:
//
//	struct A {
//	    a: String,
//	    b: Process,
//	    c: i32,
//	    d: TcpStream
//	}
//
// can be done by creating a new data message with `create_data(tag, capacity)`. `capacity` can
// be used as a hint to the host to pre-reserve the right buffer size. After a message is created,
// all the resources can be added to it with `add_*`, in this case the fields `b` & `d`. The
// returned values will be the indexes inside the message.
//
// Now the struct can be serialized for example into something like this:
//
// ["Some string" | [resource 0] | i32 value | [resource 1] ]
//
// [resource 0] & [resource 1] are just encoded as 0 and 1 u64 values, representing their index
// in the message. Now the message can be sent to another process with `send`.
//
// An important limitation here is that messages can only be worked on one at a time. If we
// called `create_data` again before sending the message, the current buffer and resources
// would be dropped.
//
// On the receiving side, first the `receive(tag)` function must be called. If `tag` has a value
// different from 0, the function will only return messages that have the specific `tag`. Once
// a message is received, we can read from its buffer or extract resources from it.
//
// This can be a bit confusing, because resources are just IDs (u64 values) themselves. But we
// still need to serialize them into different u64 values. Resources are inherently bound to a
// process and you can't access another resource just by guessing an ID from another process.
// The process of sending them around needs to be explicit.
//
// This API was designed around the idea that most guest languages will use some serialization
// library and turning resources into indexes is a way of serializing. The same is true for
// deserializing them on the receiving side, when an index needs to be turned into an actual
// resource ID.
package message

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

var (
	CallTimedOut = errors.New("call timed out")
	LinkDied     = errors.New("link died")
	ProcessDied  = errors.New("process died")
)

type ptr = unsafe.Pointer
type size = uint32

func mkptr[T any](v *T) ptr { return unsafe.Pointer(v) }

// CreateData creates a new data message.
//
// This message is intended to be modified by other functions in this namespace.
// Once `message.Send` is called, it will be sent to another process.
//
//go:wasmimport lunatic::message create_data
//go:noescape
func CreateData(tag int64, bufferCapacity uint64)

//go:wasmimport lunatic::message write_data
//go:noescape
func write_data(dataPtr ptr, dataLen size) uint32

// WriteData writes some data into the message buffer and returns how much
// data is written in bytes.
//
// Returns:
// * nil on success with number of bytes written.
// * error if there is no data message within the scratch area.
func WriteData(data []byte) (n uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.write_data error: %v", r)
		}
	}()

	n = write_data(mkptr(&data[0]), size(len(data)))
	return n, nil
}

//go:wasmimport lunatic::message read_data
//go:noescape
func read_data(dataPtr ptr, dataLen size) uint32

// ReadData reads some data from the message buffer and returns
// how many bytes were read.
//
// Returns:
// * nil on success with number of bytes read.
// * error if there is no data message within the scratch area.
func ReadData(buf []byte) (n uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.read_data error: %v", r)
		}
	}()

	n = read_data(mkptr(&buf[0]), size(len(buf)))
	return n, nil
}

//go:wasmimport lunatic::message seek_data
//go:noescape
func seek_data(index uint64)

// SeekData moves reading head of the internal message buffer.
// This is useful if you wish to read a bit of a message, decide that
// something else will handle it, `SeekData(0)` to reset the read
// position for the new receiver, and `Send` it to another process.
//
// Errors:
// * If there is no data message within the scratch area.
func SeekData(index uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.seek_data error: %v", r)
		}
	}()

	seek_data(index)
	return nil
}

//go:wasmimport lunatic::message get_tag
//go:noescape
func get_tag() int64

// GetTag returns the mssage tag or 0 if no tag was set.
//
// Returns:
// * nil if success with message tag.
// * error if there is no data message within the scratch area.
func GetTag() (tag int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.get_tag error: %v", r)
		}
	}()

	tag = get_tag()
	return tag, nil
}

//go:wasmimport lunatic::message data_size
//go:noescape
func data_size() uint64

// DataSize returns the size in bytes of the message buffer.
//
// Returns:
// * nil if success with message buffer size.
// * error if there is no data message within the scratch area.
func DataSize() (n uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.data_size error: %v", r)
		}
	}()

	n = data_size()
	return n, nil
}

//go:wasmimport lunatic::message push_tcp_stream
//go:noescape
func push_tcp_stream(streamID uint64) uint64

// PushTCPStream adds a TCP stream resource to the message that is currently
// in the scratch area and returns the new location of it.
// This will remove the TCP stream from the current process' resources.
//
// Returns:
// * nil if success with stream index.
// * error if there is no data message within the scratch area.
func PushTCPStream(streamID uint64) (index uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.push_tcp_stream error: %v", r)
		}
	}()

	index = push_tcp_stream(streamID)
	return index, nil
}

//go:wasmimport lunatic::message take_tcp_stream
//go:noescape
func take_tcp_stream(index uint64) uint64

// TakeTCPStream takes the TCP stream from the message that is currently in the scratch
// area by index, puts it into the process' resources and returns the resource ID.
//
// Returns:
// * nil if success with resource ID.
// * error if there is no data message within the scratch area.
func TakeTCPStream(index uint64) (resourceID uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.take_tcp_stream error: %v", r)
		}
	}()

	resourceID = take_tcp_stream(index)
	return resourceID, nil
}

//go:wasmimport lunatic::message push_udp_socket
//go:noescape
func push_udp_socket(socketID uint64) uint64

// PushUDPSocket adds a UDP socket resource to the message that is currently in the scratch
// area and returns the new location of it.
// This will remove the socket from the current process' resources.
//
// Returns:
// * nil if success with stream index.
// * error if there is no data message within the scratch area.
func PushUDPSocket(socketID uint64) (index uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.push_udp_socket error: %v", r)
		}
	}()

	index = push_udp_socket(socketID)
	return index, nil
}

//go:wasmimport lunatic::message take_udp_socket
//go:noescape
func take_udp_socket(index uint64) uint64

// TakeUDPSocket takes the UDP socket from the message that is currently in the scratch
// area by index, puts it into the process' resources and returns the resourceID.
//
// Returns:
// * nil if success with resource ID.
// * error if there is no data message within the scratch area.
func TakeUDPSocket(index uint64) (resourceID uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.take_udp_socket error: %v", r)
		}
	}()

	resourceID = take_udp_socket(index)
	return resourceID, nil
}

//go:wasmimport lunatic::message send
//go:noescape
func send(processID uint64) uint32

// Send sends the message to a process.
//
// There are no guarantees that the message will be received.
//
// Returns:
// * nil if successful
// * error if the processID doesn't exist or it's called before creating the next message.
func Send(processID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.send error: %v", r)
		}
	}()

	n := send(processID)
	if n != 0 {
		return fmt.Errorf("message.send error: %v", n)
	}
	return nil
}

//go:wasmimport lunatic::message send_receive_skip_search
//go:noescape
func send_receive_skip_search(processID uint64, waitOnTag int64, timeoutDuration uint64) uint32

// SendReceiveSkipSearch sends the message to a process and waits for a reply, but doesn't
// look through existing messages in the mailbox queue while waiting.
// This is an optimization that only makes sense with tagged messages.
// In a request/reply scenario, we can tag the request message with a unique tag
// and just wait on it specifically.
//
// This operation needs to be an atomic host function. If we jumped back into the guest,
// we could miss out on the incoming message before `Receive` is called.
//
// If `timeoutMillis` is not nil, the function will return on timeout expiration with
// the error `CallTimedOut`.
//
// Returns:
// * nil if message arrived.
// * ProcessDoesNotExist if the process does not exist.
// * NodeDoesNotExist if the node does not exist.
// * CallTimedOut if the call timed out.
//
// Errors:
// * If called with wrong data in the scratch area.
// * If the message contains resources.
func SendReceiveSkipSearch(processID uint64, waitOnTag int64, timeoutMillis *uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.send_receive_skip_search error: %v", r)
		}
	}()

	td := uint64(math.MaxUint64)
	if timeoutMillis != nil {
		td = *timeoutMillis
	}

	errno := send_receive_skip_search(processID, waitOnTag, td)
	switch errno {
	case 0:
		return nil
	case 9027:
		return CallTimedOut
	default:
		return fmt.Errorf("message.send_receive_skip_search unknown error %v", errno)
	}
}

//go:wasmimport lunatic::message receive
//go:noescape
func receive(tagPtr ptr, tagLen size, timeoutDuration uint64) uint32

// Receive takes the next message out of the queue or blocks until the next message is
// received if the queue is empty.
//
// If `tags` is not empty, it will block until a message is received matching any
// of the supplied tags.
//
// If `timeoutMillis` is not nil, the function will return on timeout expiration with
// the error `CallTimedOut`.
//
// Once the message is successfully received, functions like `message.ReadData()` can
// be used to extract data out of it.
//
// Returns:
// * nil if a data message arrived.
// * LinkDied if the link died.
// * ProcessDied if the process died.
// * CallTimedOut if the call timed out.
//
// Errors:
// * If there were buffer problems.
func Receive(tags []int64, timeoutMillis *uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("message.receive error: %v", r)
		}
	}()

	td := uint64(math.MaxUint64)
	if timeoutMillis != nil {
		td = *timeoutMillis
	}

	errno := receive(mkptr(&tags[0]), size(uintptr(len(tags))*unsafe.Sizeof(int64(0))), td)
	switch errno {
	case 0:
		return nil
	case 1:
		return LinkDied
	case 2:
		return ProcessDied
	case 9027:
		return CallTimedOut
	default:
		return fmt.Errorf("message.receive unknown error %v", errno)
	}
}
