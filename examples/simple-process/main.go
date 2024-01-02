// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

// simple-process is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"log"

	"github.com/gmlewis/go-lunatic/lunatic"
)

//go:wasm-module simple-process
func main() {
	log.Printf("[main process] Hello from the main process!")
	log.Printf("[main process] Spawning a child...")

	_, err := lunatic.SpawnFunc(func() {
		log.Printf("[subprocess] 👋 from spawned process!")
	})
	must(err)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
