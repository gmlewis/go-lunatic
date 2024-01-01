// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package sqlite provides the Go bindings to the lunatic::sqlite API.
package sqlite

import (
	"errors"
	"fmt"
	"unsafe"
)

type size = uint32

//go:wasmimport lunatic::sqlite open
//go:noescape
func open(pathStrPtr unsafe.Pointer, pathStrLen size, connectionIDPtr unsafe.Pointer) uint64

// Open opens a sqlite connection.
func Open(path string) (connectionID uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.open error: %v", r)
		}
	}()

	errno := open(unsafe.Pointer(&path), size(len(path)), unsafe.Pointer(&connectionID))
	switch errno {
	case 0:
		return connectionID, nil
	case 1:
		return connectionID, errors.New("path error")
	default:
		return connectionID, fmt.Errorf("sqlite.open unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::sqlite execute
//go:noescape
func execute(connID uint64, execStrPtr unsafe.Pointer, execStrLen size) uint32

// Execute executes a sqlite query.
func Execute(connID uint64, exec string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.execute error: %v", r)
		}
	}()

	errno := execute(connID, unsafe.Pointer(&exec), size(len(exec)))
	switch errno {
	case 0:
		return nil
	case 1:
		return errors.New("sqlite error")
	default:
		return fmt.Errorf("sqlite.execute unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::sqlite bind_value
//go:noescape
func bind_value(statementID uint64, bindDataPtr unsafe.Pointer, bindDataLen size)

// BindValue binds a value.
func BindValue(statementID uint64, bindData []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.bind_value error: %v", r)
		}
	}()

	bind_value(statementID, unsafe.Pointer(&bindData[0]), size(len(bindData)))
	return nil
}

//go:wasmimport lunatic::sqlite sqlite3_changes
//go:noescape
func sqlite3_changes(connID uint64) uint32

// Changes returns the sqlite change count.
func Changes(connID uint64) (changeCount uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.sqlite3_changes error: %v", r)
		}
	}()

	n := sqlite3_changes(connID)
	return n, nil
}

//go:wasmimport lunatic::sqlite statement_reset
//go:noescape
func statement_reset(statementID uint64)

// StatementReset resets a sqlite statement.
func StatementReset(statementID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.statement_reset error: %v", r)
		}
	}()

	statement_reset(statementID)
	return nil
}

//go:wasmimport lunatic::sqlite sqlite3_step
//go:noescape
func sqlite3_step(statementID uint64) uint32

// Step returns SQLITE_DONE or SQLITE_ROW depending on whether
// there's more data available or not.
func Step(statementID uint64) (status uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.sqlite3_step error: %v", r)
		}
	}()

	status = sqlite3_step(statementID)
	return status, nil
}

//go:wasmimport lunatic::sqlite sqlite3_finalize
//go:noescape
func sqlite3_finalize(statementID uint64)

// Finalize
func Finalize(statementID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.sqlite3_finalize error: %v", r)
		}
	}()

	sqlite3_finalize(statementID)
	return nil
}

//go:wasmimport lunatic::sqlite column_count
//go:noescape
func column_count(statementID uint64) uint32

// ColumnCount returns the column count.
func ColumnCount(statementID uint64) (count uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.column_count error: %v", r)
		}
	}()

	count = column_count(statementID)
	return count, nil
}

//go:wasmimport lunatic::sqlite last_error
//go:noescape
func last_error(connID uint64, opaquePtr unsafe.Pointer) uint32

// LastError returns the last error message in the provided buffer
func LastError(connID uint64, buf []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.last_error error: %v", r)
		}
	}()

	last_error(connID, unsafe.Pointer(&buf[0]))
	return nil
}

//go:wasmimport lunatic::sqlite read_column
//go:noescape
func read_column(statementID uint64, colIdx uint32, opaquePtr unsafe.Pointer) uint32

// ReadColumn reads a column at the given index.
func ReadColumn(statementID uint64, colIdx uint32, buf []byte) (n uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.read_column error: %v", r)
		}
	}()

	n = read_column(statementID, colIdx, unsafe.Pointer(&buf[0]))
	return n, nil
}

//go:wasmimport lunatic::sqlite read_row
//go:noescape
func read_row(statementID uint64, opaquePtr unsafe.Pointer) uint32

// ReadRow reads a row starting at colIdx 0.
func ReadRow(statementID uint64, buf []byte) (n uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.read_row error: %v", r)
		}
	}()

	n = read_row(statementID, unsafe.Pointer(&buf[0]))
	return n, nil
}

//go:wasmimport lunatic::sqlite column_name
//go:noescape
func column_name(statementID uint64, columnIdx uint32, opaquePtr unsafe.Pointer) uint32

// ColumnName returns the column name at the columnIdx.
func ColumnName(statementID uint64, columnIdx uint32, buf []byte) (n uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.column_name error: %v", r)
		}
	}()

	n = column_name(statementID, columnIdx, unsafe.Pointer(&buf[0]))
	return n, nil
}

//go:wasmimport lunatic::sqlite column_names
//go:noescape
func column_names(statementID uint64, opaquePtr unsafe.Pointer) uint32

// ColumnNames returns the columns names.
// TODO: Are these a vector of zero-terminated strings?
func ColumnNames(statementID uint64, buf []byte) (n uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sqlite.column_name error: %v", r)
		}
	}()

	n = column_names(statementID, unsafe.Pointer(&buf[0]))
	return n, nil
}
