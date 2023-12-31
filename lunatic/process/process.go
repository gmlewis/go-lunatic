// Package process provides access to Lunatic Process functions.
// See: https://lunatic.solutions/
package process

//go:wasmimport lunatic::process sleep_ms
//go:noescape
func sleep_ms(millis int64)

func Sleep(millis int64) {
	sleep_ms(millis)
}
