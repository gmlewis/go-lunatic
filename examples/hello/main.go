// -*- compile-command: "./build-tinygo.sh"; -*-

// hello is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import "log"

//go:wasm-module hello
func main() {
	log.Printf("Hello lunatic from Go!")
}
