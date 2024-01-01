// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package error provides the Go bindings to the lunatic::error API.
package error

import (
	"fmt"
	"unsafe"
)

type ptr = unsafe.Pointer

func mkptr[T any](v *T) ptr { return unsafe.Pointer(v) }

// StringSize returns the size of the string representation of the error `errorID`.
//
//go:wasmimport lunatic::error string_size
//go:noescape
func StringSize(errorID uint64) uint32

//go:wasmimport lunatic::error to_string
//go:noescape
func to_string(errorID uint64, errorStrPtr ptr)

// ToString returns the string representation of the error.
func ToString(errorID uint64) string {
	n := StringSize(errorID)
	buf := make([]byte, 0, n)
	to_string(errorID, mkptr(&buf[0]))
	return string(buf)
}

//go:wasmimport lunatic::error drop
//go:noescape
func drop(errorID uint64)

// Drop drops the error resource.
//
// Errors:
// * If the error ID doesn't exist.
func Drop(errorID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error.drop error: %v", r)
		}
	}()

	drop(errorID)
	return nil
}
