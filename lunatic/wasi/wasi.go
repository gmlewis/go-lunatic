// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package wasi provides the Go bindings to the lunatic::wasi API.
package wasi

import (
	"fmt"
	"unsafe"
)

type size = uint32

//go:wasmimport lunatic::wasi config_add_environment_variable
//go:noescape
func config_add_environment_variable(configID uint64, keyPtr unsafe.Pointer, keyLen size, valuePtr unsafe.Pointer, valueLen size)

// ConfigAddEnvironmentVariable adds an environment variable to a configuration.
//
// Returns:
// * nil if successful.
// * error if configID doesn't exist or `key` is an invalid UTF8 string.
func ConfigAddEnvironmentVariable(configID uint64, key, value string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("wasi.config_add_environment_variable error: %v", r)
		}
	}()

	config_add_environment_variable(configID, unsafe.Pointer(&key), size(len(key)), unsafe.Pointer(&value), size(len(value)))
	return nil
}

//go:wasmimport lunatic::wasi config_add_command_line_argument
//go:noescape
func config_add_command_line_argument(configID uint64, argumentPtr unsafe.Pointer, argumentLen size)

// ConfigAddCommandLineArgument adds a command line argument to a configuration.
//
// Returns:
// * nil if successful.
// * error if configID doesn't exist or `argument` is an invalid UTF8 string.
func ConfigAddCommandLineArgument(configID uint64, argument string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("wasi.config_add_command_line_argument error: %v", r)
		}
	}()

	config_add_command_line_argument(configID, unsafe.Pointer(&argument), size(len(argument)))
	return nil
}

//go:wasmimport lunatic::wasi config_preopen_dir
//go:noescape
func config_preopen_dir(configID uint64, dirPtr unsafe.Pointer, dirLen uint32)

// ConfigPreopenDir marks a directory as pre-opened in the configuration.
//
// Returns:
// * nil if successful.
// * error if configID doesn't exist or `dir` is an invalid UTF8 string.
func ConfigPreopenDir(configID uint64, dir string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("wasi.config_preopen_dir error: %v", r)
		}
	}()

	config_preopen_dir(configID, unsafe.Pointer(&dir), size(len(dir)))
	return nil
}
