// -*- compile-command: "./build-tinygo.sh"; -*-

// simple-process is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"log"

	"github.com/gmlewis/go-lunatic/lunatic/process"
)

//go:wasm-module simple-process
func main() {
	log.Printf("[main process] Hello from the main process!")
	log.Printf("[main process] Spawning a child...")

	if _, err := process.SpawnFunc(func() {
		log.Printf("[subprocess] 👋 from spawned process!")
	}); err != nil {
		log.Fatal(err)
	}
}
