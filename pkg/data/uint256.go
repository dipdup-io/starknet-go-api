package data

import (
	"fmt"
	"math/big"

	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	uint256low, _  = big.NewInt(0).SetString("340282366920938463463374607431768211455", 0)
	uint256high, _ = big.NewInt(0).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 0)
)

// Uint256 -
type Uint256 struct {
	low  Felt
	high Felt
}

// NewUint256 -
func NewUint256(low, high Felt) Uint256 {
	return Uint256{low, high}
}

// NewUint256FromInt -
func NewUint256FromInt(value int) Uint256 {
	low := fmt.Sprintf("0x%x", value)
	return Uint256{Felt(low), "0x0"}
}

// NewUint256FromStrings -
func NewUint256FromStrings(low, high string) Uint256 {
	return NewUint256(Felt(low), Felt(high))
}

// NewUint256FromString -
func NewUint256FromString(value string) (Uint256, error) {
	d, ok := big.NewInt(0).SetString(value, 0)
	if !ok {
		return Uint256{}, errors.Errorf("invalid uint256 string: %s", value)
	}

	highInt := big.NewInt(0).And(d, uint256high)
	lowInt := big.NewInt(0).And(d, uint256low)

	high := encoding.AddHexPrefix(highInt.Rsh(highInt, 128).Text(16))
	low := encoding.AddHexPrefix(lowInt.Text(16))
	return Uint256{Felt(low), Felt(high)}, nil
}

// Decimal -
func (uint256 Uint256) Decimal() (decimal.Decimal, error) {
	high, ok := big.NewInt(0).SetString(uint256.high.String(), 0)
	if !ok {
		return decimal.Zero, errors.Errorf("invalid high of uint256: %s", uint256.high)
	}
	low, ok := big.NewInt(0).SetString(uint256.low.String(), 0)
	if !ok {
		return decimal.Zero, errors.Errorf("invalid low of uint256: %s", uint256.high)
	}

	return decimal.NewFromBigInt(high.Lsh(high, 128).Add(high, low), 0), nil
}

// String -
func (uint256 Uint256) String() string {
	if d, err := uint256.Decimal(); err == nil {
		return d.String()
	}
	return fmt.Sprintf("low=%s high=%s", uint256.low, uint256.high)
}

// Calldata -
func (uint256 Uint256) Calldata() []string {
	return []string{
		uint256.low.Decimal().String(),
		uint256.high.Decimal().String(),
	}
}
