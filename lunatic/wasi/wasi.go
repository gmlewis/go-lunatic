// Package wasi provides the Go bindings to the lunatic::wasi API.
package wasi


// ConfigAddEnvironmentVariable
//
//go:wasmimport lunatic::wasi config_add_environment_variable
//go:noescape
func ConfigAddEnvironmentVariable(param int64 int int int int)

// ConfigAddCommandLineArgument
//
//go:wasmimport lunatic::wasi config_add_command_line_argument
//go:noescape
func ConfigAddCommandLineArgument(param int64 int int)

// ConfigPreopenDir
//
//go:wasmimport lunatic::wasi config_preopen_dir
//go:noescape
func ConfigPreopenDir(param int64 int int)
