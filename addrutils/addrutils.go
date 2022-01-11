// Package addrutils provides utilities for dealing with go-ethereum's common.Address type.
package addrutils

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

const rawZeroAddress string = "0x0000000000000000000000000000000000000000"

var ZeroAddress = common.HexToAddress(rawZeroAddress)

// IsZeroAddress returns true if the passed address
// is the "zero" address.
func IsZeroAddress(address common.Address) bool {
	return AddressEq(ZeroAddress, address)
}

func AddressEq(a, b common.Address) bool {
	return bytes.Equal(a.Bytes(), b.Bytes())
}
