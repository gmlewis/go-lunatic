// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

// spawn is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
package main

import (
	"log"

	"github.com/gmlewis/go-lunatic/lunatic/process"
)

type SomeData[T any] struct {
	Key   string
	Value T
}

//go:wasm-module spawn
func main() {
	log.Printf("[main process] Hello from the main process!")
	log.Printf("[main process] Spawning a child...")

	mailbox := process.CreateMailbox()
	localData := SomeData[int]{Key: "localData", Value: 54321}

	child, err := process.SpawnFunc(func() {
		log.Printf("[subprocess] 👋 from spawned process!")
		log.Printf("[subprocess] receiving some data...")
		log.Printf(process.Receive(mailbox))
		log.Printf("[subprocess] data within the closure:")
		log.Printf("%#v", localData)
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[main process] sending some data...")

	data := SomeData[int]{Key: "dataKey", Value: 42}
	process.Send(child, data, mailbox)
	process.SleepMS(100)
}
