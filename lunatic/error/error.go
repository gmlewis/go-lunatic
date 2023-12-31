// Package error provides the Go bindings to the lunatic::error API.
package error


// StringSize
//
//go:wasmimport lunatic::error string_size
//go:noescape
func StringSize(param int64) (int)

// ToString
//
//go:wasmimport lunatic::error to_string
//go:noescape
func ToString(param int64 int)

// Drop
//
//go:wasmimport lunatic::error drop
//go:noescape
func Drop(param int64)
