// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package version provides the Go bindings to the lunatic::version API.
package version

// Major returns the major version number.
//
//go:wasmimport lunatic::version major
//go:noescape
func Major() uint32

// Minor returns the minor version number.
//
//go:wasmimport lunatic::version minor
//go:noescape
func Minor() uint32

// Patch returns the patch version number.
//
//go:wasmimport lunatic::version patch
//go:noescape
func Patch() uint32
