// -*- compile-command: "./build-tinygo.sh"; -*-

// version is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"fmt"
	"log"

	"github.com/gmlewis/go-lunatic/lunatic/version"
)

func semverString() string {
	return fmt.Sprintf("v%v.%v.%v", version.Major(), version.Minor(), version.Patch())
}

//go:wasm-module version
func main() {
	log.Printf("Hello lunatic %v from Go!", semverString())
}
