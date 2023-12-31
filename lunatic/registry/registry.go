// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package registry provides the Go bindings to the lunatic::registry API.
package registry

// Put
//
//go:wasmimport lunatic::registry put
//go:noescape
func Put(param int int int64 int64)

// Get
//
//go:wasmimport lunatic::registry get
//go:noescape
func Get(param int int int int) (int)

// Remove
//
//go:wasmimport lunatic::registry remove
//go:noescape
func Remove(param int int)
