// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

package networking

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var (
	NotConnected = errors.New("not connected")
)

//go:wasmimport lunatic::networking udp_bind
//go:noescape
func udp_bind(addrType uint32, addrU8Ptr ptr, port, flowInfo, scopeID uint32, idU64Ptr ptr) uint32

// UDPBind creates a new UDP socket which will be bound to the specified address.
// The returned socket is ready to receive messages.
//
// Binding with a port number of 0 will request that the OS assigns a port to this socket.
// The port allocated can be queried via the `UDPLocalAddr` method.
//
// Returns:
// * nil on success with the ID of the newly-created UDP socket.
// * error with the error ID.
func UDPBind(dnsInfo DNSInfo) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_bind error: %v", r)
		}
	}()

	errno := udp_bind(dnsInfo.AddrType, mkptr(&dnsInfo.IP[0]), dnsInfo.Port, dnsInfo.FlowInfo, dnsInfo.ScopeID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	default:
		return id, fmt.Errorf("networking.udp_bind unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking drop_udp_socket
//go:noescape
func drop_udp_socket(udpSocketID uint64)

// DropUDPSocket drops the UDP socket resource.
func DropUDPSocket(udpSocketID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.drop_udp_socket error: %v", r)
		}
	}()

	drop_udp_socket(udpSocketID)
	return nil
}

//go:wasmimport lunatic::networking udp_local_addr
//go:noescape
func udp_local_addr(udpSocketID uint64, idU64Ptr ptr) uint32

// UDPLocalAddr returns the local address that this socket is bound to as
// a DNS iterator with just one element.
func UDPLocalAddr(udpSocketID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_local_addr error: %v", r)
		}
	}()

	errno := udp_local_addr(udpSocketID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	default:
		return id, fmt.Errorf("networking.udp_local_addr unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking udp_receive
//go:noescape
func udp_receive(socketID uint64, bufferPtr ptr, bufferLen size, opaquePtr ptr) uint32

// UDPReceive reads data from the connected UDP socket and writes it to the given `buf`.
// This method will fail if the socket is not connected.
func UDPReceive(socketID uint64, buf []byte) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_receive error: %v", r)
		}
	}()

	errno := udp_receive(socketID, mkptr(&buf[0]), size(cap(buf)), mkptr(&id))
	switch errno {
	case 0:
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf[0]))
		sh.Len = int(id) // override the buf slice's length to the returned results.
		return id, nil
	case 1:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.udp_receive unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking udp_receive_from
//go:noescape
func udp_receive_from(socketID uint64, bufferPtr ptr, bufferLen size, opaquePtr, dnsIterPtr ptr) uint32

// UDPReceiveFrom receives data from the UDP socket.
func UDPReceiveFrom(socketID uint64, buf []byte) (id, dnsIterID uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_receive_from error: %v", r)
		}
	}()

	errno := udp_receive_from(socketID, mkptr(&buf[0]), size(cap(buf)), mkptr(&id), mkptr(&dnsIterID))
	switch errno {
	case 0:
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf[0]))
		sh.Len = int(id) // override the buf slice's length to the returned results.
		return id, dnsIterID, nil
	default:
		return id, dnsIterID, fmt.Errorf("networking.udp_receive_from unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking udp_connect
//go:noescape
func udp_connect(addrType uint32, addrU8Ptr ptr, port, flowInfo, scopeID uint32, timeoutDuration uint64, idU64Ptr ptr) uint32

// UDPConnect connects the UDP socket to the provided dnsInfo remote address.
//
// When connected, `UDPSend` and `UDPReceive` will use the speficied address for sending and receiving messages.
// Additionally, a filter will be applied to `UDPReceiveFrom` so that it only receives messages from that same address.
//
// Returns:
// * nil on success with the ID of the newly-created UDP stream.
// * CallTimedOut if the call timed out.
// * error with the error ID.
func UDPConnect(dnsInfo DNSInfo, timeoutMillis *uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_connect error: %v", r)
		}
	}()

	td := uint64(math.MaxUint64)
	if timeoutMillis != nil {
		td = *timeoutMillis
	}

	errno := udp_connect(dnsInfo.AddrType, mkptr(&dnsInfo.IP[0]), dnsInfo.Port, dnsInfo.FlowInfo, dnsInfo.ScopeID, td, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 9027:
		return id, CallTimedOut
	default:
		return id, fmt.Errorf("networking.udp_connect unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking clone_udp_socket
//go:noescape
func clone_udp_socket(udpSocketID uint64) uint64

// CloneUDPSocket clones a UDP socket returning the ID of the clone.
func CloneUDPSocket(udpSocketID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.clone_udp_socket error: %v", r)
		}
	}()

	id = clone_udp_socket(udpSocketID)
	return id, nil
}

//go:wasmimport lunatic::networking set_udp_socket_broadcast
//go:noescape
func set_udp_socket_broadcast(udpSocketID uint64, broadcast uint32)

// SetUDPSocketBroadcast sets the broadcast state of the UDP socket.
func SetUDPSocketBroadcast(udpSocketID uint64, broadcast uint32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.set_udp_socket_broadcast error: %v", r)
		}
	}()

	set_udp_socket_broadcast(udpSocketID, broadcast)
	return nil
}

//go:wasmimport lunatic::networking get_udp_socket_broadcast
//go:noescape
func get_udp_socket_broadcast(udpSocketID uint64) int32

// GetUDPSocketBroadcast gets the current broadcast state of the UDP socket.
func GetUDPSocketBroadcast(udpSocketID uint64) (broadcast int32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.get_udp_socket_broadcast error: %v", r)
		}
	}()

	broadcast = get_udp_socket_broadcast(udpSocketID)
	return broadcast, nil
}

//go:wasmimport lunatic::networking set_udp_socket_ttl
//go:noescape
func set_udp_socket_ttl(udpSocketID uint64, ttl uint32)

// SetUDPSocketTTL sets the TTL of the UDP socket.
// This value represents the time-to-live field that is used in
// every packet sent from this socket.
func SetUDPSocketTTL(udpSocketID uint64, ttl uint32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.set_udp_socket_ttl error: %v", r)
		}
	}()

	set_udp_socket_ttl(udpSocketID, ttl)
	return nil
}

//go:wasmimport lunatic::networking get_udp_socket_ttl
//go:noescape
func get_udp_socket_ttl(udpSocketID uint64) uint32

// GetUDPSocketTTL gets the socket ttl for the UDP socket.
func GetUDPSocketTTL(udpSocketID uint64) (ttl uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.get_udp_socket_ttl error: %v", r)
		}
	}()

	ttl = get_udp_socket_ttl(udpSocketID)
	return ttl, nil
}

//go:wasmimport lunatic::networking udp_send_to
//go:noescape
func udp_send_to(socketID uint64, bufferPtr ptr, bufferLen size, addrType uint32, addrU8Ptr ptr, port, flowInfo, scopeID uint32, opaquePtr ptr) uint32

// UDPSendTo sends data on the socket to the given address.
func UDPSendTo(socketID uint64, buffer []byte, dnsInfo DNSInfo) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_send_to error: %v", r)
		}
	}()

	errno := udp_send_to(socketID, mkptr(&buffer[0]), size(len(buffer)),
		dnsInfo.AddrType, mkptr(&dnsInfo.IP[0]), dnsInfo.Port, dnsInfo.FlowInfo, dnsInfo.ScopeID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.udp_send_to unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking udp_send
//go:noescape
func udp_send(socketID uint64, bufferPtr ptr, bufferLen size, opaquePtr ptr) uint32

// UDPSend sends data on the socket to the remote address to which it is connected.
//
// The `UDPConnect` method will connect this socket to a remote address.
// This method will fail if the socket is not connected.
func UDPSend(socketID uint64, buffer []byte) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_send error: %v", r)
		}
	}()

	errno := udp_send(socketID, mkptr(&buffer[0]), size(len(buffer)), mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.udp_send unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking udp_peer_addr
//go:noescape
func udp_peer_addr(udpStreamID uint64, idU64Ptr ptr) uint32

// UDPPeerAddr returns the remote address this UDP socket is connected to, bound to a DNS
// iterator with just one element.
func UDPPeerAddr(udpSocketID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.udp_peer_addr error: %v", r)
		}
	}()

	errno := udp_peer_addr(udpSocketID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, NotConnected
	case 2:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.udp_peer_addr unknown error: %v", errno)
	}
}
