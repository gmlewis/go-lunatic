// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package registry provides the Go bindings to the lunatic::registry API.
package registry

import (
	"fmt"
	"unsafe"
)

type ptr = unsafe.Pointer
type size = uint32

func mkptr[T any](v *T) ptr { return unsafe.Pointer(v) }

//go:wasmimport lunatic::registry put
//go:noescape
func put(nameStrPtr ptr, nameStrLen size, nodeID, processID uint64)

// Put registers process with `processID` under `name`.
func Put(name string, nodeID, processID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("registry.put error: %v", r)
		}
	}()

	put(mkptr(&name), size(len(name)), nodeID, processID)
	return nil
}

//go:wasmimport lunatic::registry get
//go:noescape
func get(nameStrPtr ptr, nameStrLen size, nodeID, processID uint64) uint32

// Get looks up process under `name` and returns if it was found.
func Get(name string, nodeID, processID uint64) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("registry.get error: %v", r)
		}
	}()

	n := get(mkptr(&name), size(len(name)), nodeID, processID)
	return n == 0, nil
}

//go:wasmimport lunatic::registry remove
//go:noescape
func remove(nameStrPtr ptr, nameStrLen size)

// Remove removes the process under `name` if it exists.
func Remove(name string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("registry.remove error: %v", r)
		}
	}()

	remove(mkptr(&name), size(len(name)))
	return nil
}
