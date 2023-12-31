// Package networking provides the Go bindings to the lunatic::networking API.
package networking


// Resolve
//
//go:wasmimport lunatic::networking resolve
//go:noescape
func Resolve(param int int int64 int) (int)

// DropDNSIterator
//
//go:wasmimport lunatic::networking drop_dns_iterator
//go:noescape
func DropDNSIterator(param int64)

// ResolveNext
//
//go:wasmimport lunatic::networking resolve_next
//go:noescape
func ResolveNext(param int64 int int int int int) (int)

// TCPBind
//
//go:wasmimport lunatic::networking tcp_bind
//go:noescape
func TCPBind(param int int int int int int) (int)

// DropTCPListener
//
//go:wasmimport lunatic::networking drop_tcp_listener
//go:noescape
func DropTCPListener(param int64)

// TCPLocalAddr
//
//go:wasmimport lunatic::networking tcp_local_addr
//go:noescape
func TCPLocalAddr(param int64 int) (int)

// TCPAccept
//
//go:wasmimport lunatic::networking tcp_accept
//go:noescape
func TCPAccept(param int64 int int) (int)

// TCPConnect
//
//go:wasmimport lunatic::networking tcp_connect
//go:noescape
func TCPConnect(param int int int int int int64 int) (int)

// DropTCPStream
//
//go:wasmimport lunatic::networking drop_tcp_stream
//go:noescape
func DropTCPStream(param int64)

// CloneTCPStream
//
//go:wasmimport lunatic::networking clone_tcp_stream
//go:noescape
func CloneTCPStream(param int64) (int64)

// TCPWriteVectored
//
//go:wasmimport lunatic::networking tcp_write_vectored
//go:noescape
func TCPWriteVectored(param int64 int int int) (int)

// TCPRead
//
//go:wasmimport lunatic::networking tcp_read
//go:noescape
func TCPRead(param int64 int int int) (int)

// SetReadTimeout
//
//go:wasmimport lunatic::networking set_read_timeout
//go:noescape
func SetReadTimeout(param int64 int64)

// GetReadTimeout
//
//go:wasmimport lunatic::networking get_read_timeout
//go:noescape
func GetReadTimeout(param int64) (int64)

// SetWriteTimeout
//
//go:wasmimport lunatic::networking set_write_timeout
//go:noescape
func SetWriteTimeout(param int64 int64)

// GetWriteTimeout
//
//go:wasmimport lunatic::networking get_write_timeout
//go:noescape
func GetWriteTimeout(param int64) (int64)

// SetPeekTimeout
//
//go:wasmimport lunatic::networking set_peek_timeout
//go:noescape
func SetPeekTimeout(param int64 int64)

// GetPeekTimeout
//
//go:wasmimport lunatic::networking get_peek_timeout
//go:noescape
func GetPeekTimeout(param int64) (int64)

// TCPFlush
//
//go:wasmimport lunatic::networking tcp_flush
//go:noescape
func TCPFlush(param int64 int) (int)

// UDPBind
//
//go:wasmimport lunatic::networking udp_bind
//go:noescape
func UDPBind(param int int int int int int) (int)

// DropUDPSocket
//
//go:wasmimport lunatic::networking drop_udp_socket
//go:noescape
func DropUDPSocket(param int64)

// UDPLocalAddr
//
//go:wasmimport lunatic::networking udp_local_addr
//go:noescape
func UDPLocalAddr(param int64 int) (int)

// UDPReceive
//
//go:wasmimport lunatic::networking udp_receive
//go:noescape
func UDPReceive(param int64 int int int) (int)

// UDPReceiveFrom
//
//go:wasmimport lunatic::networking udp_receive_from
//go:noescape
func UDPReceiveFrom(param int64 int int int int) (int)

// UDPConnect
//
//go:wasmimport lunatic::networking udp_connect
//go:noescape
func UDPConnect(param int64 int int int int int int64 int) (int)

// CloneUDPSocket
//
//go:wasmimport lunatic::networking clone_udp_socket
//go:noescape
func CloneUDPSocket(param int64) (int64)

// SetUDPSocketBroadcast
//
//go:wasmimport lunatic::networking set_udp_socket_broadcast
//go:noescape
func SetUDPSocketBroadcast(param int64 int)

// GetUDPSocketBroadcast
//
//go:wasmimport lunatic::networking get_udp_socket_broadcast
//go:noescape
func GetUDPSocketBroadcast(param int64) (int)

// SetUDPSocketTTL
//
//go:wasmimport lunatic::networking set_udp_socket_ttl
//go:noescape
func SetUDPSocketTTL(param int64 int)

// GetUDPSocketTTL
//
//go:wasmimport lunatic::networking get_udp_socket_ttl
//go:noescape
func GetUDPSocketTTL(param int64) (int)

// UDPSendTo
//
//go:wasmimport lunatic::networking udp_send_to
//go:noescape
func UDPSendTo(param int64 int int int int int int int int) (int)

// UDPSend
//
//go:wasmimport lunatic::networking udp_send
//go:noescape
func UDPSend(param int64 int int int) (int)

// TCPPeerAddr
//
//go:wasmimport lunatic::networking tcp_peer_addr
//go:noescape
func TCPPeerAddr(param int64 int) (int)

// UDPPeerAddr
//
//go:wasmimport lunatic::networking udp_peer_addr
//go:noescape
func UDPPeerAddr(param int64 int) (int)
