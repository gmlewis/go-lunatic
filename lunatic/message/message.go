// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package message provides the Go bindings to the lunatic::message API.
package message

// CreateData
//
//go:wasmimport lunatic::message create_data
//go:noescape
func CreateData(param int64 int64)

// WriteData
//
//go:wasmimport lunatic::message write_data
//go:noescape
func WriteData(param int int) (int)

// ReadData
//
//go:wasmimport lunatic::message read_data
//go:noescape
func ReadData(param int int) (int)

// SeekData
//
//go:wasmimport lunatic::message seek_data
//go:noescape
func SeekData(param int64)

// GetTag
//
//go:wasmimport lunatic::message get_tag
//go:noescape
func GetTag() int64

// DataSize
//
//go:wasmimport lunatic::message data_size
//go:noescape
func DataSize() int64

// PushTCPStream
//
//go:wasmimport lunatic::message push_tcp_stream
//go:noescape
func PushTCPStream(param int64) (int64)

// TakeTCPStream
//
//go:wasmimport lunatic::message take_tcp_stream
//go:noescape
func TakeTCPStream(param int64) (int64)

// PushUDPSocket
//
//go:wasmimport lunatic::message push_udp_socket
//go:noescape
func PushUDPSocket(param int64) (int64)

// TakeUDPSocket
//
//go:wasmimport lunatic::message take_udp_socket
//go:noescape
func TakeUDPSocket(param int64) (int64)

// Send
//
//go:wasmimport lunatic::message send
//go:noescape
func Send(param int64) (int)

// SendReceiveSkipSearch
//
//go:wasmimport lunatic::message send_receive_skip_search
//go:noescape
func SendReceiveSkipSearch(param int64 int64 int64) (int)

// Receive
//
//go:wasmimport lunatic::message receive
//go:noescape
func Receive(param int int int64) (int)
