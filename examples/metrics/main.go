// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

// metrics is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
//
// This example is based on:
// https://github.com/lunatic-solutions/lunatic-rs/blob/main/examples/metrics.rs
//
// Lunatic with metrics enabled is required (enabled by default)
// To collect the metrics, prometheus feature should be also enabled
// and lunatic has to be run with --prometheus flag.
package main

import (
	"log"
	"math"

	"github.com/gmlewis/go-lunatic/lunatic/metrics"
	"github.com/gmlewis/go-lunatic/lunatic/process"
)

//go:wasm-module metrics
func main() {
	metrics.Counter("lunatic::metrics_example::counter", 42)
	for i := 0; i < 6000; i++ {
		metrics.IncrementCounter("lunatic::metrics_example::counter")
		if i%50 < 25 {
			metrics.IncrementGauge("lunatic::metrics_example::gauge", 1.0)
		} else {
			metrics.DecrementGauge("lunatic::metrics_example::gauge", 1.0)
		}
		metrics.Histogram("lunatic::metrics_example::histogram", math.Mod(float64(i), 50))
		process.SleepMS(10)
	}

	log.Printf("Done.")
}
