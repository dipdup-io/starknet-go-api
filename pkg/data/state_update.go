package data

// Nonce -
type Nonce struct {
	ContractAddress Felt `json:"contract_address"`
	Nonce           Felt `json:"nonce"`
}

// StorageDiff -
type StorageDiff struct {
	Address        Felt       `json:"address"`
	StorageEntries []KeyValue `json:"storage_entries"`
}

// KeyValue -
type KeyValue struct {
	Key   Felt `json:"key"`
	Value Felt `json:"value"`
}

// DeployedContract -
type DeployedContract struct {
	Address   Felt `json:"address"`
	ClassHash Felt `json:"class_hash"`
}

// StateDiff -
type StateDiff struct {
	StorageDiffs      map[Felt][]KeyValue `json:"storage_diffs"`
	DeclaredContracts []Felt              `json:"declared_contracts"`
	DeployedContracts []DeployedContract  `json:"deployed_contracts"`
	Nonces            map[Felt]Felt       `json:"nonces"`
}

// StateUpdate -
type StateUpdate struct {
	BlockHash Felt      `json:"block_hash"`
	NewRoot   Felt      `json:"new_root"`
	OldRoot   Felt      `json:"old_root"`
	StateDiff StateDiff `json:"state_diff"`
}
