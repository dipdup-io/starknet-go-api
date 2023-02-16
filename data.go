package starknetgoapi

import (
	stdJSON "encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Request -
type Request struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      uint64 `json:"id"`

	timeout time.Duration
}

// Response -
type Response[T any] struct {
	Version string `json:"jsonrpc"`
	Result  T      `json:"result,omitempty"`
	ID      uint64 `json:"id"`
	Error   *Error `json:"error,omitempty"`
}

// Error -
type Error struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    stdJSON.RawMessage `json:"data"`
}

// BlockRequest -
type BlockRequest struct {
	BlockNumber *uint64 `json:"block_number,omitempty"`
	BlockHash   *string `json:"block_hash,omitempty"`
}

// Nonce -
type Nonce struct {
	ContractAddress string `json:"contract_address"`
	Nonce           string `json:"nonce"`
}

// StorageDiff -
type StorageDiff struct {
	Address        string     `json:"address"`
	StorageEntries []KeyValue `json:"storage_entries"`
}

// KeyValue -
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// DeployedContract -
type DeployedContract struct {
	Address   string `json:"address"`
	ClassHash string `json:"class_hash"`
}

// StateDiff -
type StateDiff struct {
	StorageDiffs           []StateDiff        `json:"storage_diffs"`
	DeclaredContractHashes []string           `json:"declared_contract_hashes"`
	DeployedContracts      []DeployedContract `json:"deployed_contracts"`
	Nonces                 []Nonce            `json:"nonces"`
}

// Event -
type Event struct {
	FromAddress string   `json:"from_address"`
	Keys        []string `json:"keys"`
	Data        []string `json:"data"`
}

// Message -
type Message struct {
	ToAddress string   `json:"to_address"`
	Payload   []string `json:"payload"`
}

// errors
var (
	ErrInvalidBlockFilter = errors.New("Only one field of Hash, Number or String in BlockFilter should be set")
	ErrEmptyBlockFilter   = errors.New("empty BlockFilter")
)

// BlockFilter -
type BlockFilter struct {
	Hash   string
	Number uint64
	String string
}

func (bf BlockFilter) validate() error {
	hash := bf.Hash != ""
	number := bf.Number > 0
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
func (bf BlockFilter) MarshalJSON() ([]byte, error) {
	if bf.Hash != "" {
		return []byte(fmt.Sprintf(`{"block_hash": "%s"}`, bf.Hash)), nil
	}
	if bf.Number > 0 {
		return []byte(fmt.Sprintf(`{"block_number": %d}`, bf.Number)), nil
	}
	if bf.String != "" {
		return json.Marshal(bf.String)
	}
	return nil, ErrEmptyBlockFilter
}

// CallRequest -
type CallRequest struct {
	ContractAddress    string   `json:"contract_address"`
	EntrypointSelector string   `json:"entry_point_selector"`
	Calldata           []string `json:"calldata"`
}
