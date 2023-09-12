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

// DeclaredClass -
type DeclaredClass struct {
	ClassHash         Felt `json:"class_hash"`
	CompiledClassHash Felt `json:"compiled_class_hash"`
}

// StateDiff -
type StateDiff struct {
	StorageDiffs         map[Felt][]KeyValue `json:"storage_diffs"`
	DeclaredClasses      []DeclaredClass     `json:"declared_classes"`
	ReplacedClasses      []ReplacedClass     `json:"replaced_classes"`
	OldDeclaredContracts []Felt              `json:"old_declared_contracts"`
	DeployedContracts    []DeployedContract  `json:"deployed_contracts"`
	Nonces               map[Felt]Felt       `json:"nonces"`
}

type ReplacedClass struct {
	Address   Felt `json:"address"`
	ClassHash Felt `json:"class_hash"`
}

// StateUpdate -
type StateUpdate struct {
	BlockHash Felt      `json:"block_hash"`
	NewRoot   Felt      `json:"new_root"`
	OldRoot   Felt      `json:"old_root"`
	StateDiff StateDiff `json:"state_diff"`
}

// StateDiffRpc -
type StateDiffRpc struct {
	DeclaredClasses        []DeclaredClass    `json:"declared_classes"`
	ReplacedClasses        []ReplacedClass    `json:"replaced_classes"`
	DeclaredContractHashes []Felt             `json:"declared_contract_hashes"`
	DeployedContracts      []DeployedContract `json:"deployed_contracts"`
	Nonces                 []Nonce            `json:"nonces"`
	StorageDiffs           []StorageDiff      `json:"storage_diffs"`
}

// ToStateDiff -
func (sdr StateDiffRpc) ToStateDiff() StateDiff {
	sd := StateDiff{
		StorageDiffs:         make(map[Felt][]KeyValue),
		Nonces:               make(map[Felt]Felt),
		DeclaredClasses:      sdr.DeclaredClasses,
		ReplacedClasses:      sdr.ReplacedClasses,
		OldDeclaredContracts: sdr.DeclaredContractHashes,
		DeployedContracts:    sdr.DeployedContracts,
	}

	for i := range sdr.Nonces {
		sd.Nonces[sdr.Nonces[i].ContractAddress] = sdr.Nonces[i].Nonce
	}

	for i := range sdr.StorageDiffs {
		sd.StorageDiffs[sdr.StorageDiffs[i].Address] = sdr.StorageDiffs[i].StorageEntries
	}

	return sd
}

// StateUpdateRpc -
type StateUpdateRpc struct {
	BlockHash Felt         `json:"block_hash"`
	NewRoot   Felt         `json:"new_root"`
	OldRoot   Felt         `json:"old_root"`
	StateDiff StateDiffRpc `json:"state_diff"`
}

// ToStateUpdate -
func (sur StateUpdateRpc) ToStateUpdate() StateUpdate {
	return StateUpdate{
		BlockHash: sur.BlockHash,
		NewRoot:   sur.NewRoot,
		OldRoot:   sur.OldRoot,
		StateDiff: sur.StateDiff.ToStateDiff(),
	}
}
