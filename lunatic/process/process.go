// Package process provides the Go bindings to the lunatic::process API.
package process


// CompileModule
//
//go:wasmimport lunatic::process compile_module
//go:noescape
func CompileModule(param int int int) (int)

// DropModule
//
//go:wasmimport lunatic::process drop_module
//go:noescape
func DropModule(param int64)

// CreateConfig
//
//go:wasmimport lunatic::process create_config
//go:noescape
func CreateConfig() int64

// DropConfig
//
//go:wasmimport lunatic::process drop_config
//go:noescape
func DropConfig(param int64)

// ConfigSetMaxMemory
//
//go:wasmimport lunatic::process config_set_max_memory
//go:noescape
func ConfigSetMaxMemory(param int64 int64)

// ConfigGetMaxMemory
//
//go:wasmimport lunatic::process config_get_max_memory
//go:noescape
func ConfigGetMaxMemory(param int64) (int64)

// ConfigSetMaxFuel
//
//go:wasmimport lunatic::process config_set_max_fuel
//go:noescape
func ConfigSetMaxFuel(param int64 int64)

// ConfigGetMaxFuel
//
//go:wasmimport lunatic::process config_get_max_fuel
//go:noescape
func ConfigGetMaxFuel(param int64) (int64)

// ConfigCanCompileModules
//
//go:wasmimport lunatic::process config_can_compile_modules
//go:noescape
func ConfigCanCompileModules(param int64) (int)

// ConfigSetCanCompileModules
//
//go:wasmimport lunatic::process config_set_can_compile_modules
//go:noescape
func ConfigSetCanCompileModules(param int64 int)

// ConfigCanCreateConfigs
//
//go:wasmimport lunatic::process config_can_create_configs
//go:noescape
func ConfigCanCreateConfigs(param int64) (int)

// ConfigSetCanCreateConfigs
//
//go:wasmimport lunatic::process config_set_can_create_configs
//go:noescape
func ConfigSetCanCreateConfigs(param int64 int)

// ConfigCanSpawnProcesses
//
//go:wasmimport lunatic::process config_can_spawn_processes
//go:noescape
func ConfigCanSpawnProcesses(param int64) (int)

// ConfigSetCanSpawnProcesses
//
//go:wasmimport lunatic::process config_set_can_spawn_processes
//go:noescape
func ConfigSetCanSpawnProcesses(param int64 int)

// Spawn
//
//go:wasmimport lunatic::process spawn
//go:noescape
func Spawn(param int64 int64 int64 int int int int int) (int)

// SleepMS
//
//go:wasmimport lunatic::process sleep_ms
//go:noescape
func SleepMS(param int64)

// DieWhenLinkDies
//
//go:wasmimport lunatic::process die_when_link_dies
//go:noescape
func DieWhenLinkDies(param int)

// ProcessID
//
//go:wasmimport lunatic::process process_id
//go:noescape
func ProcessID() int64

// Link
//
//go:wasmimport lunatic::process link
//go:noescape
func Link(param int64 int64)

// Unlink
//
//go:wasmimport lunatic::process unlink
//go:noescape
func Unlink(param int64)

// Kill
//
//go:wasmimport lunatic::process kill
//go:noescape
func Kill(param int64)

// Exists
//
//go:wasmimport lunatic::process exists
//go:noescape
func Exists(param int64) (int)
