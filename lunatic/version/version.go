// Package version provides the Go bindings to the lunatic::version API.
package version

// Major
//
//go:wasmimport lunatic::version major
//go:noescape
func Major() int

// Minor
//
//go:wasmimport lunatic::version minor
//go:noescape
func Minor() int

// Patch
//
//go:wasmimport lunatic::version patch
//go:noescape
func Patch() int
