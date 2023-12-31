// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package distributed provides the Go bindings to the lunatic::distributed API.
package distributed

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var (
	CallTimedOut        = errors.New("call timed out")
	ModuleDoesNotExist  = errors.New("module does not exist")
	NodeConnectionError = errors.New("node connection error")
	NodeDoesNotExist    = errors.New("node does not exist")
	ProcessDoesNotExist = errors.New("process does not exist")
)

type size = uint32

// NodesCount returns the number of registered nodes.
//
//go:wasmimport lunatic::distributed nodes_count
//go:noescape
func NodesCount() uint32

//go:wasmimport lunatic::distributed get_nodes
//go:noescape
func get_nodes(buf unsafe.Pointer, bufLen size) uint32

// GetNodes copies node IDs into the `ids` slice which must have
// enough capacity to hold the results.
// It returns the number of nodes copied.
//
// Returns:
// * error if ids slice is not large enough to hold all the nodes IDs.
func GetNodes(ids []uint64) (err error) {
	var n uint32
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("distributed.get_nodes error: ids.cap=%v, n=%v: %v", cap(ids), n, r)
		}
	}()

	n = get_nodes(unsafe.Pointer(&ids[0]), size(len(ids)))
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&ids))
	sh.Len = int(n) // override the slice's length to the returned results.
	return nil
}

// NodeID returns the ID of the node that the current process is running on.
//
//go:wasmimport lunatic::distributed node_id
//go:noescape
func NodeID() uint64

// ModuleID returns the ID of the module that the current process is spawned from.
//
//go:wasmimport lunatic::distributed module_id
//go:noescape
func ModuleID() uint64

//go:wasmimport lunatic::distributed spawn
//go:noescape
func spawn(nodeID uint64, configID int64, moduleID uint64,
	funcStrPtr unsafe.Pointer, funcStrLen size, paramsPtr unsafe.Pointer, paramsLen size, idPtr unsafe.Pointer) uint32

// Spawn spawns a new process using the passed-in function inside a module
// as the entry point. The process is spawned on a node with ID `nodeID`.
//
// If `configID` is 0, the same config is used as in the process calling
// this function.
//
// The function arguments are passed as a slice of params of any size integer.
//
// Returns:
// * nil on success with the `id` of the newly-created process.
// * NodeDoesNotExist if the node does not exist.
// * ModuleDoesNotExist if the module does not exist.
// * NodeConnectionError if a node connection error occurred.
// * error for unsupported params types (must be any size integer).
//
// Errors:
// * If the function string is not a valid UTF8 string.
// * If the params array is in the wrong format.
// * If any memory outside this guest heap space is referenced.
func Spawn(nodeID uint64, configID int64, moduleID uint64, funcStr string, params []any) (id uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("distributed.spawn error: %v", r)
		}
	}()

	paramsBytes := make([]byte, 18*len(params))
	copyBytes := func(index int, v uint64) {
		src := unsafe.Slice((*byte)(unsafe.Pointer(&v)), 17)
		for i := 0; i < 17; i++ {
			paramsBytes[index*18+i] = src[i]
		}
	}

	for i, param := range params {
		i32func := func(v uint64) {
			paramsBytes[i*18] = 0x7F // i32
			copyBytes(i, v)
		}
		i64func := func(v uint64) {
			paramsBytes[i*18] = 0x7E // i64
			copyBytes(i, v)
		}

		switch t := param.(type) {
		case int8:
			i32func(uint64(t))
		case int16:
			i32func(uint64(t))
		case int:
			i32func(uint64(t))
		case int32:
			i32func(uint64(t))
		case uint8:
			i32func(uint64(t))
		case uint16:
			i32func(uint64(t))
		case uint:
			i32func(uint64(t))
		case uint32:
			i32func(uint64(t))
		case int64:
			i64func(uint64(t))
		case uint64:
			i64func(uint64(t))
		case uintptr:
			i64func(uint64(t))
		// case i128, u128:  // https://github.com/golang/go/issues/9455#issuecomment-74165846
		default:
			return id, fmt.Errorf("distributed.spawn params[%v] = %T, expected integer", i, param)
		}
	}

	errno := spawn(nodeID, configID, moduleID, unsafe.Pointer(&funcStr), size(len(funcStr)),
		unsafe.Pointer(&paramsBytes[0]), size(len(paramsBytes)), unsafe.Pointer(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, NodeDoesNotExist
	case 2:
		return id, ModuleDoesNotExist
	case 9027:
		return id, NodeConnectionError
	default:
		return id, fmt.Errorf("distributed.spawn unknown error %v", errno)
	}
}

//go:wasmimport lunatic::distributed send
//go:noescape
func send(nodeID, processID uint64) uint32

// Send sends the message in scratch area to a process running on a node with ID `nodeID`.
//
// There are no guarantees that the message will be received.
//
// Returns:
// * nil if message sent.
// * ProcessDoesNotExist if the process does not exist.
// * NodeDoesNotExist if the node does not exist.
// * NodeConnectionError if a node connection error occurred.
//
// Errors:
// * If called before creating the next message.
// * If the message contains resources.
func Send(nodeID, processID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("distributed.send error: %v", r)
		}
	}()

	errno := send(nodeID, processID)
	switch errno {
	case 0:
		return nil
	case 1:
		return ProcessDoesNotExist
	case 2:
		return NodeDoesNotExist
	case 9027:
		return NodeConnectionError
	default:
		return fmt.Errorf("distributed.send unknown error %v", errno)
	}
}

//go:wasmimport lunatic::distributed send_receive_skip_search
//go:noescape
func send_receive_skip_search(nodeID, processID uint64, waitOnTag int64, timeoutDuration uint64) uint32

// SendReceiveSkipSearch sends the message to a process on a node with ID `nodeID` and waits for a reply,
// but doesn't look through existing messages in the mailbox queue while waiting.
// This is an optimization that only makes sense with tagged messages.
// In a request/reply scenario we can tag the request message with a unique tag and just wait on it specifically.
//
// This operation needs to be an atomic host function. If we jumped back into the guest we could
// miss out on the incoming message before `receive` is called.
//
// If timeoutMillis is not nil, the function will return on timeout expiration with the error CallTimedOut.
//
// Returns:
// * nil if message arrived.
// * ProcessDoesNotExist if the process does not exist.
// * NodeDoesNotExist if the node does not exist.
// * CallTimedOut.
//
// Errors:
// * If called with wrong data in the scratch area.
// * If the message contains resources.
func SendReceiveSkipSearch(nodeID, processID uint64, waitOnTag int64, timeoutMillis *uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("distributed.send_receive_skip_search error: %v", r)
		}
	}()

	td := uint64(math.MaxUint64)
	if timeoutMillis != nil {
		td = *timeoutMillis
	}

	errno := send_receive_skip_search(nodeID, processID, waitOnTag, td)
	switch errno {
	case 0:
		return nil
	case 1:
		return ProcessDoesNotExist
	case 2:
		return NodeDoesNotExist
	case 9027:
		return CallTimedOut
	default:
		return fmt.Errorf("distributed.send_receive_skip_search unknown error %v", errno)
	}
}
