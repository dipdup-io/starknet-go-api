package api

import (
	"time"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
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
	Version string      `json:"jsonrpc"`
	Result  T           `json:"result,omitempty"`
	ID      uint64      `json:"id"`
	Error   *data.Error `json:"error,omitempty"`
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

// CallRequest -
type CallRequest struct {
	ContractAddress    string   `json:"contract_address"`
	EntrypointSelector string   `json:"entry_point_selector"`
	Calldata           []string `json:"calldata"`
}
