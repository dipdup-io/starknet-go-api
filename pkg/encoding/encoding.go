package encoding

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// entrypoint names
const (
	DefaultEntrypoint         = "__default__"
	ConstructorEntrypoint     = "constructor"
	DefaultL1Entrypoint       = "__l1_default__"
	ExecuteEntrypoint         = "__execute__"
	TransferEntrypoint        = "transfer"
	ValidateEntrypoint        = "__validate__"
	ValidateDeclareEntrypoint = "__validate_declare__"
	ValidateDeployEntrypoint  = "__validate_deploy__"
	ChangeModulesEntrypoint   = "changeModules"

	DefaultEntrypoitSelector = 0x0
)

var (
	ExecuteEntrypointSelector         = MustDecodeHex("0x15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad")
	ConstructorSelector               = MustDecodeHex("0x28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194")
	ValidateDeclareEntrypointSelector = MustDecodeHex("0x289da278a8dc833409cabfdad1581e8e7d40e42dcaed693fa4008dcdb4963b3")
	ChangeModuleEntrypointSelector    = MustDecodeHex("0x3ffada7235f48d4811be030385f19e6d50e2cfa368ded42f1892666f834e407")
)

// Keccak - A variant of eth-keccak that computes a value that fits in a StarkNet field element.
func Keccak(data []byte) []byte {
	bytes := crypto.Keccak256(data)
	bytes[0] &= 0x3
	return bytes
}

// GetSelectorFromName -
func GetSelectorFromName(name string) string {
	bytes := Keccak([]byte(name))
	return strings.TrimLeft(hex.EncodeToString(bytes), "0")
}

// GetSelectorWithPrefixFromName -
func GetSelectorWithPrefixFromName(name string) string {
	return AddHexPrefix(GetSelectorFromName(name))
}

// TrimHex - trims prefix '0x' if it exists and all padding left zeroes.
func TrimHex(val string) string {
	return strings.TrimLeft(strings.TrimPrefix(val, "0x"), "0")
}

// AddHexPrefix -
func AddHexPrefix(s string) string {
	return "0x" + s
}

// MustDecodeHex -
func MustDecodeHex(s string) []byte {
	s = strings.TrimPrefix(s, "0x")
	if len(s)%2 == 1 {
		s = "0" + s
	}
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// EncodeHex -
func EncodeHex(data []byte) string {
	return AddHexPrefix(hex.EncodeToString(data))
}

// TrimmedHex -
func TrimmedHex(data []byte) string {
	s := hex.EncodeToString(data)
	return strings.TrimLeft(s, "0")
}

// DecimalFromHex -
func DecimalFromHex(s string) decimal.Decimal {
	if s == "" {
		return decimal.Zero
	}
	i, _ := new(big.Int).SetString(s, 0)
	return decimal.NewFromBigInt(i, 0)
}
