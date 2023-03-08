package encoding

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
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

	DefaultEntrypoitSelector = 0x0
)

var (
	ExecuteEntrypointSelector = MustDecodeHex("0x15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad")
	ConstructorSelector       = MustDecodeHex("0x28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194")
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

// TrimHex - trims prefix '0x' if it exists and all padding left zeroes.
func TrimHex(val string) string {
	return strings.TrimLeft(strings.TrimPrefix(val, "0x"), "0")
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
	return "0x" + hex.EncodeToString(data)
}

// TrimmedHex -
func TrimmedHex(data []byte) string {
	s := hex.EncodeToString(data)
	return strings.TrimLeft(s, "0")
}
