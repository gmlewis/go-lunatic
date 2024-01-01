// -*- compile-command: "./build-tinygo.sh"; -*-

// sleep is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"fmt"

	"github.com/gmlewis/go-lunatic/lunatic/process"
)

//go:wasm-module sleep
func main() {
	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		process.SleepMS(200)
		fmt.Println(i)
	}

	fmt.Println("Done.")
}
