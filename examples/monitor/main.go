// -*- compile-command: "./build-go.sh && ./run.sh"; -*-

// monitor is a simple example of creating a WASM module that can
// be run by lunatic. See: https://lunatic.solutions/
//
// This example is based on:
// https://github.com/lunatic-solutions/lunatic-rs/blob/main/examples/monitor.rs
package main

import (
	"log"

	"github.com/gmlewis/go-lunatic/lunatic"
)

//go:wasm-module monitor
func main() { // mailbox any) {
	args, err := lunatic.Args()
	must(err)
	if len(args) != 1 {
		log.Fatalf("expected mailbox name, but args=%+v", args)
	}
	log.Printf("args=%+v", args)

	// mailboxName := args[0]

	/*
	   let mailbox = mailbox.monitorable();

	   let process = Process::spawn(mailbox.this(), child_process);
	   mailbox.monitor(process);

	   loop {
	       match mailbox.receive() {
	           MessageSignal::Message(msg) => {
	               println!("{msg}");
	           }
	           MessageSignal::Signal(ProcessDiedSignal(id)) => {
	               println!("Process {id} died");
	               break;
	           }
	       }
	   }
	*/

	log.Printf("Done.")
}

func childProcess(parent, mailbox any) {
	log.Printf("childProcess: parent=%[1]T=%[1]v, mailbox=%[2]T=%[2]v", parent, mailbox)
	// parent.send("Hello".to_string());
	// lunatic::sleep(Duration::from_secs(3));
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
