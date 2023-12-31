// Package sqlite provides the Go bindings to the lunatic::sqlite API.
package sqlite


// Open
//
//go:wasmimport lunatic::sqlite open
//go:noescape
func Open(param int int int) (int64)

// Execute
//
//go:wasmimport lunatic::sqlite execute
//go:noescape
func Execute(param int64 int int) (int)

// BindValue
//
//go:wasmimport lunatic::sqlite bind_value
//go:noescape
func BindValue(param int64 int int)

// Sqlite3Changes
//
//go:wasmimport lunatic::sqlite sqlite3_changes
//go:noescape
func Sqlite3Changes(param int64)(result int)

// StatementReset
//
//go:wasmimport lunatic::sqlite statement_reset
//go:noescape
func StatementReset(param int64)

// Sqlite3Step
//
//go:wasmimport lunatic::sqlite sqlite3_step
//go:noescape
func Sqlite3Step(param int64) (int)

// Sqlite3Finalize
//
//go:wasmimport lunatic::sqlite sqlite3_finalize
//go:noescape
func Sqlite3Finalize(param int64)

// ColumnCount
//
//go:wasmimport lunatic::sqlite column_count
//go:noescape
func ColumnCount(param int64) (int)

// LastError
//
//go:wasmimport lunatic::sqlite last_error
//go:noescape
func LastError(param int64 int) (int)

// ReadColumn
//
//go:wasmimport lunatic::sqlite read_column
//go:noescape
func ReadColumn(param int64 int int) (int)

// ReadRow
//
//go:wasmimport lunatic::sqlite read_row
//go:noescape
func ReadRow(param int64 int) (int)

// ColumnName
//
//go:wasmimport lunatic::sqlite column_name
//go:noescape
func ColumnName(param int64 int int) (int)

// ColumnNames
//
//go:wasmimport lunatic::sqlite column_names
//go:noescape
func ColumnNames(param int64 int) (int)
