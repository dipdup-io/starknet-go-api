package data

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
	"github.com/shopspring/decimal"
)

// Felt -
type Felt string

// NewFeltFromBytes -
func NewFeltFromBytes(hash []byte) Felt {
	return Felt(encoding.EncodeHex(hash))
}

// NewFromAsciiString -
func NewFromAsciiString(s string) Felt {
	return Felt("0x" + hex.EncodeToString([]byte(s)))
}

// String -
func (f Felt) String() string {
	return string(f)
}

// Decimal -
func (f Felt) Decimal() decimal.Decimal {
	return encoding.DecimalFromHex(f.String())
}

// Uint64 -
func (f Felt) Uint64() (uint64, error) {
	if f == "" {
		return 0, nil
	}
	return strconv.ParseUint(f.String(), 0, 64)
}

// Bytes -
func (f Felt) Bytes() []byte {
	return encoding.MustDecodeHex(f.String())
}

// ToAsciiString -
func (f Felt) ToAsciiString() (string, error) {
	s := strings.TrimPrefix(f.String(), "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
