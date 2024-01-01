// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

package networking

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"unsafe"
)

var (
	CallTimedOut = errors.New("call timed out")
)

type ptr = unsafe.Pointer
type size = uint32

func mkptr[T any](v *T) ptr { return unsafe.Pointer(v) }

//go:wasmimport lunatic::networking resolve
//go:noescape
func resolve(nameStrPtr ptr, nameStrLen size, timeoutDuration uint64, idU64Ptr ptr) uint32

// Resolve performs a DNS resolution. The returned iterator may not actually yield any values
// depending on the outcome of any resolution performed.
//
// If `timeoutMillis` is not nil, the function will return on timeout expiration with
// the error `CallTimedOut`.
//
// Returns:
// * nil on success with the ID of the newly created DNS iterator.
// * error with the ID of the error.
func Resolve(name string, timeoutMillis *uint64) (id uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.resolve error: %v", r)
		}
	}()

	td := uint64(math.MaxUint64)
	if timeoutMillis != nil {
		td = *timeoutMillis
	}

	log.Printf("Resolve calling resolve(0x%x,%v,%v,0x%x)", mkptr(&name), size(len(name)), td, mkptr(&id))
	errno := resolve(mkptr(&name), size(len(name)), td, mkptr(&id))
	log.Printf("errno=%v, id=%v", errno, id)
	switch errno {
	case 0:
		return id, nil
	case 1:
		return id, errors.New("error id")
	case 9027:
		return id, CallTimedOut
	default:
		return id, fmt.Errorf("networking.resolve unknown error: %v", errno)
	}
}

//go:wasmimport lunatic::networking drop_dns_iterator
//go:noescape
func drop_dns_iterator(dnsIterID uint64)

// DropDNSIterator drops the DNS iterator resource.
func DropDNSIterator(dnsIterID uint64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.drop_dns_iterator error: %v", r)
		}
	}()

	drop_dns_iterator(dnsIterID)
	return nil
}

//go:wasmimport lunatic::networking resolve_next
//go:noescape
func resolve_next(dnsIterID uint64, addrTypeU32Ptr, addrU8Ptr, portU16Ptr, flowInfoU32Ptr, scopeIDU32Ptr ptr) uint32

// ResolveNext takes the next socket address from the DNS iterator and returns it.
// When the iterator is exhausted, (nil, nil) is returned.
func ResolveNext(dnsIterID uint64) (info *DNSInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("networking.resolve_next error: %v", r)
		}
	}()

	var (
		addrType uint32
		ip       net.IP = make([]byte, 16)
		port     uint32
		flowInfo uint32
		scopeID  uint32
	)

	dnsInfo := &DNSInfo{
		IP: make([]byte, 0, 16),
	}

	// For some reason, the rust binding uses a U16 pointer here for port whereas
	// the rest of lunatic uses U32 for ports.
	log.Printf("ResolveNext calling resolve_next(%v,%v,%v,%v,%v,%v)", dnsIterID, mkptr(&addrType), mkptr(&ip), mkptr(&port), mkptr(&flowInfo), mkptr(&scopeID))
	n := resolve_next(dnsIterID, mkptr(&dnsInfo.AddrType), mkptr(&dnsInfo.IP), mkptr(&dnsInfo.Port), mkptr(&dnsInfo.FlowInfo), mkptr(&dnsInfo.ScopeID))
	log.Printf("got n=%v, dnsInfo=%#v", n, dnsInfo)
	switch n {
	case 0:
		// dnsInfo.Port = uint32(port)
		return dnsInfo, nil
	case 1:
		return nil, nil
	default:
		return nil, fmt.Errorf("networking.resolve_next unknown error: %v", n)
	}
}
