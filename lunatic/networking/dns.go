// -*- compile-command: "GOOS=wasip1 GOARCH=wasm go test ./..."; -*-

package networking

import (
	"errors"
	"fmt"
	"math"
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
	nameBytes := []byte(name)

	errno := resolve(mkptr(&nameBytes[0]), size(len(name)), td, mkptr(&id))
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

	dnsInfo := &DNSInfo{
		IP: make([]byte, 16),
	}

	n := resolve_next(dnsIterID, mkptr(&dnsInfo.AddrType), mkptr(&dnsInfo.IP[0]), mkptr(&dnsInfo.Port), mkptr(&dnsInfo.FlowInfo), mkptr(&dnsInfo.ScopeID))
	switch n {
	case 0:
		return dnsInfo, nil
	case 1:
		return nil, nil
	default:
		return nil, fmt.Errorf("networking.resolve_next unknown error: %v", n)
	}
}
