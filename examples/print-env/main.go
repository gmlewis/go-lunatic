// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

// print-env is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
//
// This example is based on:
// https://github.com/lunatic-solutions/lunatic-rs/blob/main/examples/print_env.rs
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/gmlewis/go-lunatic/lunatic"
)

//go:wasm-module print-env
func main() {
	args, err := lunatic.Args()
	must(err)
	fmt.Printf("args: %+v\n", args)

	envVars := os.Environ()
	sort.Strings(envVars)
	for _, line := range envVars {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("%v\n", line)
			continue
		}
		key, value := parts[0], parts[1]
		fmt.Printf("%v: %v\n", key, value)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
