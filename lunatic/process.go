// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

package lunatic

import (
	"github.com/gmlewis/go-lunatic/lunatic/process"
)

// SpawnFunc is a helper to spawn a new function in Go.
func SpawnFunc(fn func()) (processID uint32, err error) {
	funcName := "__lunatic_bootstrap"

	processID, err = process.Spawn(0, -1, -1, funcName, nil)
	if err != nil {
		return processID, err
	}

	// TODO: finish this.
	// nodeID := distributed.NodeID()
	// message.Send()
	return processID, nil
}
