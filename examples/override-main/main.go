// -*- compile-command: "./build-tinygo.sh actual_main && ./run.sh"; -*-

// override-main is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
//
// This example shows how the mailbox parameter passed by Lunatic can be used within Go.
//
// Note that the following tools must be installed for this example to work:
//
// * `wasm2wat`
// * `wat2wasm`
//
// Both tools are available from: https://github.com/WebAssembly/wabt
//
// Until `//go:wasmexport` is implemented (https://github.com/golang/go/issues/42372),
// this hack only works with TinyGo.
package main

import (
	"log"
	"unsafe"
)

// main is used to make the initial "main.wasm" file.
// But afterward, the "main.wasm" file is post-processed
// to make the entry point be the "actual_main" function
// below so that we can pass in the mailbox arg from Lunatic.
func main() { actual_main(nil) }

// actual_main is the actual entry-point to this wasm module from Lunatic.
//
//go:wasm-module override-main
//export actual_main
func actual_main(mailbox unsafe.Pointer) {
	log.Printf("mailbox: %v", mailbox)
	log.Printf("Hello lunatic from Go!")
}
