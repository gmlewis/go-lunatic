// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package metrics provides the Go bindings to the lunatic::metrics API.
package metrics

import (
	"fmt"
	"unsafe"
)

type ptr = unsafe.Pointer
type size = uint32

func mkptr[T any](v *T) ptr { return unsafe.Pointer(v) }

//go:wasmimport lunatic::metrics counter
//go:noescape
func counter(nameStrPtr ptr, nameStrLen size, value uint64)

// Counter sets a counter.
func Counter(name string, value uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("metrics.counter error: %v", r)
		}
	}()

	counter(mkptr(&name), size(len(name)), value)
	return nil
}

//go:wasmimport lunatic::metrics increment_counter
//go:noescape
func increment_counter(nameStrPtr ptr, nameStrLen size)

// IncrementCounter increments a counter.
func IncrementCounter(name string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("metrics.increment_counter error: %v", r)
		}
	}()

	increment_counter(mkptr(&name), size(len(name)))
	return nil
}

//go:wasmimport lunatic::metrics gauge
//go:noescape
func gauge(nameStrPtr ptr, nameStrLen size, value float64)

// Gauge sets a guage.
func Gauge(name string, value float64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("metrics.gauge error: %v", r)
		}
	}()

	gauge(mkptr(&name), size(len(name)), value)
	return nil
}

//go:wasmimport lunatic::metrics increment_gauge
//go:noescape
func increment_gauge(nameStrPtr ptr, nameStrLen size, value float64)

// IncrementGauge increments a gauge.
func IncrementGauge(name string, value float64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("metrics.increment_gauge error: %v", r)
		}
	}()

	increment_gauge(mkptr(&name), size(len(name)), value)
	return nil
}

//go:wasmimport lunatic::metrics decrement_gauge
//go:noescape
func decrement_gauge(nameStrPtr ptr, nameStrLen size, value float64)

// DecrementGauge decrements a gauge.
func DecrementGauge(name string, value float64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("metrics.decrement_gauge error: %v", r)
		}
	}()

	decrement_gauge(mkptr(&name), size(len(name)), value)
	return nil
}

//go:wasmimport lunatic::metrics histogram
//go:noescape
func histogram(nameStrPtr ptr, nameStrLen size, value float64)

// Histogram sets a histogram.
func Histogram(name string, value float64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("metrics.histogram error: %v", r)
		}
	}()

	histogram(mkptr(&name), size(len(name)), value)
	return nil
}
