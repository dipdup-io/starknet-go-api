package data

import (
	"bytes"
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
	if len(f) == 0 {
		return nil
	}
	data := encoding.MustDecodeHex(f.String())
	if len(data) < AddressBytesLength {
		padding := bytes.Repeat([]byte{0}, AddressBytesLength-len(data))
		data = append(padding, data...)
	}
	return data
}

// ToAsciiString -
func (f Felt) ToAsciiString() string {
	s := strings.TrimPrefix(f.String(), "0x")
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return f.String()
	}
	return string(b)
}
