// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

package networking

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

//go:wasmimport lunatic::networking tcp_bind
//go:noescape
func tcp_bind(addrType uint32, addrU8Ptr ptr, port, flowInfo, scopeID uint32, idU64Ptr ptr) uint32

// TCPBind creates a new TCP listener which will be bound to the specified address.
// The returned listener is ready to accept connections.
//
// Binding with a port number of 0 will request that the OS assigns a port to this listener.
// The port allocated can be queried via the `TCPLocalAddr` method.
//
// Returns:
// * nil on success with the ID of the newly-created TCP listener.
// * error with the error ID.
func TCPBind(dnsInfo DNSInfo) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_bind error: %v", r)
		}
	}()

	errno := tcp_bind(dnsInfo.AddrType, mkptr(&dnsInfo.IP[0]), dnsInfo.Port, dnsInfo.FlowInfo, dnsInfo.ScopeID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	default:
		return id, fmt.Errorf("networking.tcp_bind unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking drop_tcp_listener
//go:noescape
func drop_tcp_listener(tcpListenerID uint64)

// DropTCPListener drops the TCP listener resource.
func DropTCPListener(tcpListenerID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.drop_tcp_listener error: %v", r)
		}
	}()

	drop_tcp_listener(tcpListenerID)
	return nil
}

//go:wasmimport lunatic::networking tcp_local_addr
//go:noescape
func tcp_local_addr(tcpListenerID uint64, idU64Ptr ptr) uint32

// TCPLocalAddr returns the local address that this listener is bound to as
// a DNS iterator with just one element.
func TCPLocalAddr(tcpListenerID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_local_addr error: %v", r)
		}
	}()

	errno := tcp_local_addr(tcpListenerID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	default:
		return id, fmt.Errorf("networking.tcp_local_addr unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking tcp_accept
//go:noescape
func tcp_accept(listenerID uint64, idU64Ptr ptr, socketAddrIDPtr ptr) uint32

// TCPAccept returns the ID of the newly-created TCP stream and the peer address
// as a DNS iterator with just one element.
func TCPAccept(listenerID uint64) (id, dnsIterID uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_accept error: %v", r)
		}
	}()

	errno := tcp_accept(listenerID, mkptr(&id), mkptr(&dnsIterID))
	switch errno {
	case 0:
		return id, dnsIterID, nil
	default:
		return id, dnsIterID, fmt.Errorf("networking.tcp_accept unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking tcp_connect
//go:noescape
func tcp_connect(addrType uint32, addrU8Ptr ptr, port, flowInfo, scopeID uint32, timeoutDuration uint64, idU64Ptr ptr) uint32

// TCPConnect connects to the provided dnsInfo.
//
// Returns:
// * nil on success with the ID of the newly-created TCP stream.
// * CallTimedOut if the call timed out.
// * error with the error ID.
func TCPConnect(dnsInfo DNSInfo, timeoutMillis *uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_connect error: %v", r)
		}
	}()

	td := uint64(math.MaxUint64)
	if timeoutMillis != nil {
		td = *timeoutMillis
	}

	errno := tcp_connect(dnsInfo.AddrType, mkptr(&dnsInfo.IP[0]), dnsInfo.Port, dnsInfo.FlowInfo, dnsInfo.ScopeID, td, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 9027:
		return id, CallTimedOut
	default:
		return id, fmt.Errorf("networking.tcp_connect unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking drop_tcp_stream
//go:noescape
func drop_tcp_stream(tcpStreamID uint64)

// DropTCPStream drops the TCP stream resource.
func DropTCPStream(tcpStreamID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.drop_tcp_stream error: %v", r)
		}
	}()

	drop_tcp_stream(tcpStreamID)
	return nil
}

//go:wasmimport lunatic::networking clone_tcp_stream
//go:noescape
func clone_tcp_stream(tcpStreamID uint64) uint64

// CloneTCPStream clones a TCP stream returning the ID of the clone.
func CloneTCPStream(tcpStreamID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.clone_tcp_stream error: %v", r)
		}
	}()

	id = clone_tcp_stream(tcpStreamID)
	return id, nil
}

//go:wasmimport lunatic::networking tcp_write_vectored
//go:noescape
func tcp_write_vectored(streamID uint64, ciovecArrayPtr ptr, ciovecArrayLen size, opaquePtr ptr) uint32

// TCPWriteVectored gathers data from the vector buffers and writes them to the stream.
func TCPWriteVectored(streamID uint64, buf []byte) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_write_vectored error: %v", r)
		}
	}()

	errno := tcp_write_vectored(streamID, mkptr(&buf[0]), size(len(buf)), mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.tcp_write_vectored unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking tcp_read
//go:noescape
func tcp_read(streamID uint64, bufferPtr ptr, bufferLen size, opaquePtr ptr) uint32

// TCPRead reads data from the TCP stream and writes it into `buf`.
//
// If no data was read within the specified timeout duration, then CallTimedOut is returned.
func TCPRead(streamID uint64, buf []byte) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_read error: %v", r)
		}
	}()

	errno := tcp_read(streamID, mkptr(&buf[0]), size(cap(buf)), mkptr(&id))
	switch errno {
	case 0:
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf[0]))
		sh.Len = int(id) // override the buf slice's length to the returned results.
		return id, nil
	case 1:
		return id, errors.New("error id")
	case 9027:
		return id, CallTimedOut
	default:
		return id, fmt.Errorf("networking.tcp_read unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking set_read_timeout
//go:noescape
func set_read_timeout(streamID, duration uint64)

// SetReadTimeout sets the new value for read timeout for the TCP stream.
func SetReadTimeout(streamID, timeoutMillis uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.set_read_timeout error: %v", r)
		}
	}()

	set_read_timeout(streamID, timeoutMillis)
	return nil
}

//go:wasmimport lunatic::networking get_read_timeout
//go:noescape
func get_read_timeout(streamID uint64) uint64

// GetReadTimeout gets the read timeout for the TCP stream.
func GetReadTimeout(streamID uint64) (timeoutMillis uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.get_read_timeout error: %v", r)
		}
	}()

	timeoutMillis = get_read_timeout(streamID)
	return timeoutMillis, nil
}

//go:wasmimport lunatic::networking set_write_timeout
//go:noescape
func set_write_timeout(streamID, duration uint64)

// SetWriteTimeout sets the new value for write timeout for the TCP stream.
func SetWriteTimeout(streamID, timeoutMillis uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.set_write_timeout error: %v", r)
		}
	}()

	set_write_timeout(streamID, timeoutMillis)
	return nil
}

//go:wasmimport lunatic::networking get_write_timeout
//go:noescape
func get_write_timeout(streamID uint64) uint64

// GetWriteTimeout gets the value for the write timeout for the TCP stream.
func GetWriteTimeout(streamID uint64) (timeoutMillis uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.get_write_timeout error: %v", r)
		}
	}()

	timeoutMillis = get_write_timeout(streamID)
	return timeoutMillis, nil
}

//go:wasmimport lunatic::networking set_peek_timeout
//go:noescape
func set_peek_timeout(streamID, duration uint64)

// SetPeekTimeout sets the new value for peek timeout for the TCP stream.
func SetPeekTimeout(streamID, timeoutMillis uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.set_peek_timeout error: %v", r)
		}
	}()

	set_peek_timeout(streamID, timeoutMillis)
	return nil
}

//go:wasmimport lunatic::networking get_peek_timeout
//go:noescape
func get_peek_timeout(streamID uint64) uint64

// GetPeekTimeout gets the value for the peek timeout for the TCP stream.
func GetPeekTimeout(streamID uint64) (timeoutMillis uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.get_peek_timeout error: %v", r)
		}
	}()

	timeoutMillis = get_peek_timeout(streamID)
	return timeoutMillis, nil
}

//go:wasmimport lunatic::networking tcp_flush
//go:noescape
func tcp_flush(streamID uint64, errorIDPtr ptr) uint32

// TCPFlush flushes this output stream, ensuring that all buffered contents
// reach their destination.
func TCPFlush(streamID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_flush error: %v", r)
		}
	}()

	errno := tcp_flush(streamID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.tcp_flush unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking tcp_peer_addr
//go:noescape
func tcp_peer_addr(tcpStreamID uint64, idU64Ptr ptr) uint32

// TCPPeerAddr returns the remote address this TCP socket is connected to, bound to a DNS
// iterator with just one element.
func TCPPeerAddr(tcpStreamID uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.tcp_peer_addr error: %v", r)
		}
	}()

	errno := tcp_peer_addr(tcpStreamID, mkptr(&id))
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, errors.New("error id")
	default:
		return id, fmt.Errorf("networking.tcp_peer_addr unknown error: %v", errno)
	}
}
