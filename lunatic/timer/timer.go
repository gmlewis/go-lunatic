// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package timer provides the Go bindings to the lunatic::timer API.
package timer

import "fmt"

//go:wasmimport lunatic::timer send_after
//go:noescape
func send_after(processID uint64, delayMillis uint64) uint64

// SendAfter sends the message to a process after a delay.
//
// There are no guarantees that the message will be received.
//
// Returns:
// * nil if successful with message ID.
// * error if the processID doesn't exist or if called before creating the next message.
func SendAfter(processID uint64, delayMillis uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("timer.send_after error: %v", r)
		}
	}()

	id = send_after(processID, delayMillis)
	return id, nil
}

//go:wasmimport lunatic::timer cancel_timer
//go:noescape
func cancel_timer(timerID uint64) uint32

// CancelTimer cancels the specified timer.
//
// Returns:
// * true if the `timerID` was found.
// * false if no timer was found with `timerID`, because:
//   - timer had expired
//   - timer already had been canceled
//   - timerID never corresponded to a timer.
func CancelTimer(timerID uint64) bool {
	n := cancel_timer(timerID)
	return n == 1
}
