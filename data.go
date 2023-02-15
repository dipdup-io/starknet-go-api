package starknetgoapi

import (
	stdJSON "encoding/json"
	"time"
)

const defaultJSONRPC = "2.0"
const latest = "latest"
const pending = "pending"

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

// Transaction -
type Transaction struct {
	Type                string   `json:"type"`
	Version             string   `json:"version"`
	TransactionHash     string   `json:"transaction_hash"`
	MaxFee              string   `json:"max_fee,omitempty"`
	ClassHash           string   `json:"class_hash,omitempty"`
	Nonce               string   `json:"nonce,omitempty"`
	SenderAddress       string   `json:"sender_address,omitempty"`
	ContractAddressSalt string   `json:"contract_address_salt,omitempty"`
	Calldata            []string `json:"calldata,omitempty"`
	Signature           []string `json:"signature,omitempty"`
	ConstructorCalldata []string `json:"constructor_calldata,omitempty"`
}

type AutoGenerated struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Type            string `json:"type"`
		TransactionHash string `json:"transaction_hash"`
		Version         string `json:"version"`
	} `json:"result"`
	ID int `json:"id"`
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
