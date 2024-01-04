// -*- compile-command: "go test ./..."; -*-

// override-main is a hack that is used to enable parameters to be
// passed from Lunatic to a TinyGo-generated main.wasm file.
//
// It performs the following steps:
//
// * convert wasmFile to wat (using `wasm2wat`)
// * modify the wat file to pass the parameter
// * convert the modified wat file back to wasm.
//
// Obviously, this is a hack and will hopefully be easier in time
// so that this hack can be deleted.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	tmpFile = flag.String("tmp", "out.wat", "Temporary wat filename")
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 || *tmpFile == "" {
		log.Fatalf("usage: override-main [-tmp out.wat] main.wasm actual_main")
	}

	wasmFile, entryPoint := flag.Arg(0), flag.Arg(1)
	cmd := fmt.Sprintf("wasm2wat '%v' > '%v'", wasmFile, *tmpFile)
	log.Printf("Running: %v", cmd)
	buf, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Fatalf("command failed: %v\n%s", err, buf)
	}

	buf, err = os.ReadFile(*tmpFile)
	if err != nil {
		log.Fatal(err)
	}

	out, err := transformWat(string(buf), entryPoint)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(*tmpFile, []byte(out), 0644); err != nil {
		log.Fatal(err)
	}

	cmd = fmt.Sprintf("wat2wasm -o '%v' '%v'", wasmFile, *tmpFile)
	log.Printf("Running: %v", cmd)
	buf, err = exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Fatalf("command failed: %v\n%s", err, buf)
	}

	log.Printf("Done.")
}

func transformWat(in, entryPoint string) (string, error) {
	// TODO: write this.
	return in, nil
}
