// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

// Package networking provides the Go bindings to the lunatic::networking API.
package networking

import "net"

// DNSInfo represents v4 or v6 DNS address info.
type DNSInfo struct {
	AddrType uint32 // '4' or '6'
	IP       net.IP // v4 or v6 IP address.
	Port     uint32
	// the following are only used for ipv6 connections:
	FlowInfo uint32
	ScopeID  uint32
}
