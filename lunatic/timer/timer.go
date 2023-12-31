// Package timer provides the Go bindings to the lunatic::timer API.
package timer


// SendAfter
//
//go:wasmimport lunatic::timer send_after
//go:noescape
func SendAfter(param int64 int64) (int64)

// CancelTimer
//
//go:wasmimport lunatic::timer cancel_timer
//go:noescape
func CancelTimer(param int64) (int)
