// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package process provides the Go bindings to the lunatic::process API.
package process

import (
	"errors"
	"fmt"
	"unsafe"
)

var (
	ModuleDoesNotExist  = errors.New("module does not exist")
	NodeConnectionError = errors.New("node connection error")
	NodeDoesNotExist    = errors.New("node does not exist")
	PermissionDenied    = errors.New("permission denied")
)

type ptr = unsafe.Pointer
type size = uint32

func mkptr[T any](v *T) ptr { return unsafe.Pointer(v) }

//go:wasmimport lunatic::process compile_module
//go:noescape
func compile_module(moduleDataPtr ptr, moduleDataLen size, idPtr ptr) int32

// CompileModule compiles a new WebAssembly module.
//
// The `Spawn` function can be used to spawn new processes from the module.
// Module compilation can be a CPU-intensive task.
//
// Returns:
// * nil on success. The ID of the newly-created module is also returned.
// * PermissionDenied if the process doesn't have permission to compile modules.
// * a wrapped wrror ID
func CompileModule(moduleData string) (id uint32, err error) {
	errno := compile_module(mkptr(&moduleData), size(len(moduleData)), mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, fmt.Errorf("process.compile_module error id: %v", id)
	case -1:
		return 0, PermissionDenied
	default:
		return 0, fmt.Errorf("process.compile_module unknown error %v", errno)
	}
}

//go:wasmimport lunatic::process drop_module
//go:noescape
func drop_module(moduleID uint64)

// DropModule drops the module from resources.
//
// Errors:
// * If the module ID does not exist.
func DropModule(moduleID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.drop_module error: %v", r)
		}
	}()

	drop_module(moduleID)
	return nil
}

//go:wasmimport lunatic::process create_config
//go:noescape
func create_config() int64

// CreateConfig creates a new configuration with all permissions denied.
//
// There is no memory or fuel limit set on the newly-created configuration.
//
// Returns:
// * ID of newly-created configuration in case of success.
// * PermissionDenied in case the process doesn't have permission to create new configurations.
func CreateConfig() (id int64, err error) {
	switch v := create_config(); v {
	case -1:
		return 0, PermissionDenied
	default:
		return v, nil
	}
}

//go:wasmimport lunatic::process drop_config
//go:noescape
func drop_config(configID uint64)

// DropConfig drops the configuration from resources.
//
// Returns:
// * Error if config ID doesn't exist.
func DropConfig(configID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.drop_config error: %v", r)
		}
	}()

	drop_config(configID)
	return nil
}

//go:wasmimport lunatic::process config_set_max_memory
//go:noescape
func config_set_max_memory(configID, maxMemory uint64)

// ConfigSetMaxMemory sets the memory limit on a configuration.
//
// Returns:
// * Error if config ID doesn't exist or if maxMemory is bigger than the platform maximum.
func ConfigSetMaxMemory(configID, maxMemory uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_set_max_memory error: %v", r)
		}
	}()

	config_set_max_memory(configID, maxMemory)
	return nil
}

//go:wasmimport lunatic::process config_get_max_memory
//go:noescape
func config_get_max_memory(configID uint64) uint64

// ConfigGetMaxMemory returns the memory limit of a configuration.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigGetMaxMemory(configID uint64) (maxMemory uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_get_max_memory error: %v", r)
		}
	}()

	n := config_get_max_memory(configID)
	return n, nil
}

//go:wasmimport lunatic::process config_set_max_fuel
//go:noescape
func config_set_max_fuel(configID, maxFuel uint64)

// ConfigSetMaxFuel sets the fuel limit on a configuration.
//
// A value of 0 indicates no fuel limit.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigSetMaxFuel(configID, maxFuel uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_set_max_fuel error: %v", r)
		}
	}()

	config_set_max_fuel(configID, maxFuel)
	return nil
}

//go:wasmimport lunatic::process config_get_max_fuel
//go:noescape
func config_get_max_fuel(configID uint64) uint64

// ConfigGetMaxFuel returns the fuel limit of a configuration.
//
// A value of 0 indicates no fuel limit.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigGetMaxFuel(configID uint64) (maxFuel uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_get_max_fuel error: %v", r)
		}
	}()

	n := config_get_max_fuel(configID)
	return n, nil
}

//go:wasmimport lunatic::process config_can_compile_modules
//go:noescape
func config_can_compile_modules(configID uint64) uint32

// ConfigCanCompileModules returns whether processes spawned from this
// configuration can compile Wasm modules.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigCanCompileModules(configID uint64) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_can_compile_modules error: %v", r)
		}
	}()

	n := config_can_compile_modules(configID)
	return n == 1, nil
}

//go:wasmimport lunatic::process config_set_can_compile_modules
//go:noescape
func config_set_can_compile_modules(configID uint64, can uint32)

// ConfigSetCanCompileModules sets whether processes spawned from this
// configuration will be able to compile Wasm modules.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigSetCanCompileModules(configID uint64, ok bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_set_can_compile_modules error: %v", r)
		}
	}()

	var can uint32
	if ok {
		can = 1
	}
	config_set_can_compile_modules(configID, can)
	return nil
}

//go:wasmimport lunatic::process config_can_create_configs
//go:noescape
func config_can_create_configs(configID uint64) uint32

// ConfigCanCreateConfigs returns whether processes spawned from this
// configuration can create other configurations.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigCanCreateConfigs(configID uint64) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_can_create_configs error: %v", r)
		}
	}()

	n := config_can_create_configs(configID)
	return n == 1, nil
}

//go:wasmimport lunatic::process config_set_can_create_configs
//go:noescape
func config_set_can_create_configs(configID uint64, can uint32)

// ConfigSetCanCreateConfigs sets whether processes spawned from this
// configuration will be able to create other configurations.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigSetCanCreateConfigs(configID uint64, ok bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_set_can_create_configs error: %v", r)
		}
	}()

	var can uint32
	if ok {
		can = 1
	}
	config_set_can_create_configs(configID, can)
	return nil
}

//go:wasmimport lunatic::process config_can_spawn_processes
//go:noescape
func config_can_spawn_processes(configID uint64) int32

// ConfigCanSpawnProcesses returns whether processes spawned from this
// configuration can spawn sub-processes.
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigCanSpawnProcesses(configID uint64) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_can_spawn_processes error: %v", r)
		}
	}()

	n := config_can_spawn_processes(configID)
	return n == 1, nil
}

//go:wasmimport lunatic::process config_set_can_spawn_processes
//go:noescape
func config_set_can_spawn_processes(configID uint64, can uint32)

// ConfigSetCanSpawnProcesses sets whether processes spawned from this
// configuration will be able to spawn sub-processes
//
// Returns:
// * Error if config ID doesn't exist.
func ConfigSetCanSpawnProcesses(configID uint64, ok bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.config_set_can_spawn_processes error: %v", r)
		}
	}()

	var can uint32
	if ok {
		can = 1
	}
	config_set_can_spawn_processes(configID, can)
	return nil
}

//go:wasmimport lunatic::process spawn
//go:noescape
func spawn(link, configID, moduleID int64, funcStrPtr ptr, funcStrLen size,
	paramsPtr ptr, paramsLen size, idPtr ptr) uint32

// SpawnFunc is a helper to spawn a new function in Go.
func SpawnFunc(fn func()) (id uint32, err error) {
	funcName := "__lunatic_bootstrap"

	id, err = Spawn(0, -1, -1, funcName, nil)
	if err != nil {
		return id, err
	}

	// TODO: complete this.
	// message.Send()
	return id, nil
}

// Spawn spawns a new process using the passed-in function inside a module as the entry point.
//
// If `link` is not 0, it will link the child and parent processes. The value of `link` will
// be used as the link-tag for the child. This means that if the child panics, the parent
// is going to get a signal back with the value used as the tag.
//
// If `configID` or `moduleID` have the value -1, the same module/config is used as in the
// process calling this function.
//
// The function arguments are passed as a slice of params of any size integer.
//
// Returns:
// * nil on success with the `id` of the newly-created process.
// * ModuleDoesNotExist if the module does not exist.
// * error for unsupported params types (must be any size integer).
//
// Errors:
// * If the function string is not a valid UTF8 string.
// * If the params array is in the wrong format.
// * If any memory outside this guest heap space is referenced.
func Spawn(link, configID, moduleID int64, funcStr string, params []any) (id uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.spawn error: %v", r)
		}
	}()

	paramsBytes := make([]byte, 18*len(params))
	copyBytes := func(index int, v uint64) {
		src := unsafe.Slice((*byte)(unsafe.Pointer(&v)), 17)
		for i := 0; i < 17; i++ {
			paramsBytes[index*18+i] = src[i]
		}
	}

	for i, param := range params {
		i32func := func(v uint64) {
			paramsBytes[i*18] = 0x7F // i32
			copyBytes(i, v)
		}
		i64func := func(v uint64) {
			paramsBytes[i*18] = 0x7E // i64
			copyBytes(i, v)
		}

		switch t := param.(type) {
		case int8:
			i32func(uint64(t))
		case int16:
			i32func(uint64(t))
		case int:
			i32func(uint64(t))
		case int32:
			i32func(uint64(t))
		case uint8:
			i32func(uint64(t))
		case uint16:
			i32func(uint64(t))
		case uint:
			i32func(uint64(t))
		case uint32:
			i32func(uint64(t))
		case int64:
			i64func(uint64(t))
		case uint64:
			i64func(uint64(t))
		case uintptr:
			i64func(uint64(t))
		// case i128, u128:  // https://github.com/golang/go/issues/9455#issuecomment-74165846
		default:
			return id, fmt.Errorf("params[%v] = %T, expected integer", i, param)
		}
	}

	errno := spawn(link, configID, moduleID, mkptr(&funcStr), size(len(funcStr)),
		mkptr(&paramsBytes[0]), size(len(paramsBytes)), mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, NodeDoesNotExist
	case 2:
		return id, ModuleDoesNotExist
	case 9027:
		return id, NodeConnectionError
	default:
		return id, fmt.Errorf("unknown error %v", errno)
	}
}

// SleepMS suspends this process for `millis` milliseconds.
//
// Returns:
// * Error if config ID doesn't exist.
//
//go:wasmimport lunatic::process sleep_ms
//go:noescape
func SleepMS(millis uint64)

//go:wasmimport lunatic::process die_when_link_dies
//go:noescape
func die_when_link_dies(trap uint32)

// DieWhenLinkDies defines what happens to this process if one of the linked processes
// notifies us that it died.
// `die` true means that the process will die and notify all linked processes of its death.
// `die` false means that the received signal will be turned into a signal message and put into
// this process' mailbox.
//
// The default behavior for a newly-spawned process is to die.
func DieWhenLinkDies(die bool) {
	var trap uint32
	if die {
		trap = 1
	}
	die_when_link_dies(trap)
}

// ProcessID returns the ID of the process currently running.
//
//go:wasmimport lunatic::process process_id
//go:noescape
func ProcessID() uint64

//go:wasmimport lunatic::process link
//go:noescape
func link(tag int64, processID uint64)

// Link links the current process to `processID`. This is not an atomic operation. Either of
// the two processes could fail before processing the `Link` signal and may not notify the other.
//
// Returns:
// * Error if process ID doesn't exist.
func Link(tag int64, processID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.link error: %v", r)
		}
	}()

	link(tag, processID)
	return nil
}

//go:wasmimport lunatic::process unlink
//go:noescape
func unlink(processID uint64)

// Unlink unlinks the current process from `processID`. This is not an atomic operation.
//
// Returns:
// * Error if process ID doesn't exist.
func Unlink(processID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.unlink error: %v", r)
		}
	}()

	unlink(processID)
	return nil
}

//go:wasmimport lunatic::process kill
//go:noescape
func kill(processID uint64)

// Kill sends a kill signal to `processID`.
//
// Returns:
// * Error if process ID doesn't exist.
func Kill(processID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("process.kill error: %v", r)
		}
	}()

	kill(processID)
	return nil
}

//go:wasmimport lunatic::process exists
//go:noescape
func exists(processID uint64) int32

// Exists returns whether the `processID` exists.
//
// Returns:
// * Error if process ID doesn't exist.
func Exists(processID uint64) bool {
	n := exists(processID)
	return n != 0
}
