package data

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
	StorageDiffs      map[string][]KeyValue `json:"storage_diffs"`
	DeclaredContracts []string              `json:"declared_contracts"`
	DeployedContracts []DeployedContract    `json:"deployed_contracts"`
	Nonces            map[string]string     `json:"nonces"`
}

// StateUpdate -
type StateUpdate struct {
	BlockHash string    `json:"block_hash"`
	NewRoot   string    `json:"new_root"`
	OldRoot   string    `json:"old_root"`
	StateDiff StateDiff `json:"state_diff"`
}
