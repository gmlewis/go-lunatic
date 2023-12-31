// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package metrics provides the Go bindings to the lunatic::metrics API.
package metrics

// Counter
//
//go:wasmimport lunatic::metrics counter
//go:noescape
func Counter(param int int int64)

// IncrementCounter
//
//go:wasmimport lunatic::metrics increment_counter
//go:noescape
func IncrementCounter(param int int)

// Gauge
//
//go:wasmimport lunatic::metrics gauge
//go:noescape
func Gauge(param int int f64)

// IncrementGauge
//
//go:wasmimport lunatic::metrics increment_gauge
//go:noescape
func IncrementGauge(param int int f64)

// DecrementGauge
//
//go:wasmimport lunatic::metrics decrement_gauge
//go:noescape
func DecrementGauge(param int int f64)

// Histogram
//
//go:wasmimport lunatic::metrics histogram
//go:noescape
func Histogram(param int int f64)
