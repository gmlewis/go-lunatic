// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

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
	// var timeoutMillis uint64 = 10000
	dnsIterID, err := networking.Resolve("google.com:80", nil) // &timeoutMillis)
	must(err)

	for {
		dnsInfo, err := networking.ResolveNext(dnsIterID)
		must(err)
		if dnsInfo == nil {
			break
		}
		fmt.Printf("dns info: %#v\n", *dnsInfo)
	}

	log.Printf("Done.")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
