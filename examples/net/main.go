// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

// net is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"fmt"
	"log"

	"github.com/gmlewis/go-lunatic/lunatic/networking"
)

const uri = "google.com:80"

//go:wasm-module net
func main() {
	dnsIterID, err := networking.Resolve(uri, nil)
	must(err)

	var count int
	for {
		dnsInfo, err := networking.ResolveNext(dnsIterID)
		must(err)
		if dnsInfo == nil {
			break
		}
		count++
		fmt.Printf("'%v' %v: dns address type %v: %v port %v\n", uri, count, dnsInfo.AddrType, dnsInfo.IP, dnsInfo.Port)
	}

	log.Printf("Done.")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
