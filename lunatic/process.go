// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

package lunatic

import (
	"errors"
	"unsafe"

	"github.com/gmlewis/go-lunatic/lunatic/process"
)

type size = uint32
type errno = uint32
type uintptr32 = uint32

//go:wasmimport wasi_snapshot_preview1 args_get
//go:noescape
func args_get(argv, argvBuf unsafe.Pointer) errno

//go:wasmimport wasi_snapshot_preview1 args_sizes_get
//go:noescape
func args_sizes_get(argc, argvBufLen unsafe.Pointer) errno

// Args returns the arguments that were passed to the Process' main (aka "_start") function.
func Args() ([]string, error) {
	// From: https://tip.golang.org/src/runtime/os_wasip1.go
	var argc size
	var argvBufLen size
	if args_sizes_get(unsafe.Pointer(&argc), unsafe.Pointer(&argvBufLen)) != 0 {
		return nil, errors.New("args_sizes_get failed")
	}

	argslice := make([]string, argc)
	if argc > 0 {
		argv := make([]uintptr32, argc)
		argvBuf := make([]byte, argvBufLen)
		if args_get(unsafe.Pointer(&argv[0]), unsafe.Pointer(&argvBuf[0])) != 0 {
			return nil, errors.New("args_get failed")
		}

		for i := range argslice {
			start := argv[i] - uintptr32(uintptr(unsafe.Pointer(&argvBuf[0])))
			end := start
			for argvBuf[end] != 0 {
				end++
			}
			argslice[i] = string(argvBuf[start:end])
		}
	}

	return argslice, nil
}

// SpawnFunc is a helper to spawn a new function in Go.
func SpawnFunc(fn func()) (processID uint32, err error) {
	funcName := "__lunatic_bootstrap"

	processID, err = process.Spawn(0, -1, -1, funcName, nil)
	if err != nil {
		return processID, err
	}

	// TODO: finish this.
	// nodeID := distributed.NodeID()
	// message.Send()
	return processID, nil
}
