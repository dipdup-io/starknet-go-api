package data

import (
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

// errors
var (
	ErrInvalidBlockFilter = errors.New("Only one field of Hash, Number or String in BlockFilter should be set")
	ErrEmptyBlockFilter   = errors.New("empty BlockFilter")
)

// BlockID -
type BlockID struct {
	Hash   string
	Number *uint64
	String string
}

// Validate -
func (bf BlockID) Validate() error {
	hash := bf.Hash != ""
	number := bf.Number != nil
	str := bf.String != ""
	switch {
	case hash && number:
		return ErrInvalidBlockFilter
	case hash && str:
		return ErrInvalidBlockFilter
	case number && str:
		return ErrInvalidBlockFilter
	case !hash && !number && !str:
		return ErrEmptyBlockFilter
	}

	if bf.String != Latest && bf.String != Pending && bf.String != "" {
		return errors.Errorf("String field in block filter has to be 'latest' or 'pending', but '%s' was passed", bf.String)
	}
	return nil
}

// MarshalJSON -
func (bf BlockID) MarshalJSON() ([]byte, error) {
	if bf.Hash != "" {
		return []byte(fmt.Sprintf(`{"block_hash": "%s"}`, bf.Hash)), nil
	}
	if bf.Number != nil {
		return []byte(fmt.Sprintf(`{"block_number": %d}`, *bf.Number)), nil
	}
	if bf.String != "" {
		return json.Marshal(bf.String)
	}
	return nil, ErrEmptyBlockFilter
}

// GetArg -
func (bf BlockID) GetArg() (string, string) {
	if bf.Hash != "" {
		return "blockHash", bf.Hash
	}
	if bf.Number != nil {
		return "blockNumber", strconv.FormatUint(*bf.Number, 10)
	}
	if bf.String != "" {
		return "blockNumber", bf.String
	}
	return "", ""
}
