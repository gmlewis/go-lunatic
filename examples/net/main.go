// -*- compile-command: "./build-go.sh"; -*-

// net is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"fmt"
	"log"

	"github.com/gmlewis/go-lunatic/lunatic/networking"
)

//go:wasm-module net
func main() {
	address, err := networking.Resolve("google.com:80", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("address: %v\n", address)
}
