package data

import (
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

// Invoke -
type Invoke struct {
	MaxFee                    Felt     `json:"max_fee"`
	Nonce                     Felt     `json:"nonce"`
	ContractAddress           Felt     `json:"contract_address"`
	EntrypointSelector        Felt     `json:"entry_point_selector"`
	SenderAddress             Felt     `json:"sender_address"`
	ChainId                   Felt     `json:"chain_id"`
	FeeDataAvailabilityMode   Felt     `json:"fee_data_availability_mode"`
	NonceDataAvailabilityMode Felt     `json:"nonce_data_availability_mode"`
	Tip                       Felt     `json:"tip"`
	Signature                 []string `json:"signature"`
	Calldata                  []string `json:"calldata"`
	AccountDeploymentData     []Felt   `json:"account_deployment_data"`
	PayMasterData             []Felt   `json:"paymaster_data"`
}

func (i Invoke) toMap(version Felt) map[string]any {
	data := map[string]any{
		"max_fee":          i.MaxFee,
		"nonce":            i.Nonce,
		"contract_address": i.ContractAddress,
		"signature":        i.Signature,
		"calldata":         i.Calldata,
	}

	switch version {
	case Version0:
		data["entry_point_selector"] = i.EntrypointSelector
	case Version1:
		data["sender_address"] = i.SenderAddress
	case Version3:
		data["chain_id"] = i.ChainId
		data["fee_data_availability_mode"] = i.FeeDataAvailabilityMode
		data["nonce_data_availability_mode"] = i.NonceDataAvailabilityMode
		data["tip"] = i.Tip
		data["account_deployment_data"] = i.AccountDeploymentData
		data["paymaster_data"] = i.PayMasterData
	}

	return data
}

// Declare -
type Declare struct {
	MaxFee                    Felt     `json:"max_fee"`
	Nonce                     Felt     `json:"nonce"`
	SenderAddress             Felt     `json:"sender_address"`
	ContractAddress           Felt     `json:"contract_address"`
	Signature                 []string `json:"signature"`
	ContractClass             *Class   `json:"contract_class,omitempty"`
	ClassHash                 Felt     `json:"class_hash,omitempty"`
	CompiledClassHash         Felt     `json:"compiled_class_hash,omitempty"`
	AccountDeploymentData     []Felt   `json:"account_deployment_data"`
	ChainId                   Felt     `json:"chain_id"`
	FeeDataAvailabilityMode   Felt     `json:"fee_data_availability_mode"`
	NonceDataAvailabilityMode Felt     `json:"nonce_data_availability_mode"`
	PayMasterData             []Felt   `json:"paymaster_data"`
	Tip                       Felt     `json:"tip"`
}

func (d Declare) toMap(version Felt) map[string]any {
	data := map[string]any{
		"max_fee":        d.MaxFee,
		"nonce":          d.Nonce,
		"sender_address": d.SenderAddress,
		"signature":      d.Signature,
		"contract_class": d.ContractClass,
	}

	switch version {
	case Version0, Version1, Version2:
		return data
	case Version3:
		data["chain_id"] = d.ChainId
		data["fee_data_availability_mode"] = d.FeeDataAvailabilityMode
		data["nonce_data_availability_mode"] = d.NonceDataAvailabilityMode
		data["tip"] = d.Tip
		data["account_deployment_data"] = d.AccountDeploymentData
		data["paymaster_data"] = d.PayMasterData
	}
	return data
}

// Deploy -
type Deploy struct {
	ContractAddressSalt string   `json:"contract_address_salt"`
	ConstructorCalldata []string `json:"constructor_calldata"`
	ClassHash           Felt     `json:"class_hash,omitempty"`
	ContractClass       Class    `json:"contract_class,omitempty"`
	ContractAddress     Felt     `json:"contract_address"`
}

func (d Deploy) toMap() map[string]any {
	return map[string]any{
		"contract_address_salt": d.ContractAddressSalt,
		"constructor_calldata":  d.ConstructorCalldata,
		"contract_class":        d.ContractClass,
	}
}

// DeployAccount -
type DeployAccount struct {
	MaxFee                    Felt     `json:"max_fee"`
	Nonce                     Felt     `json:"nonce"`
	ContractAddress           Felt     `json:"contract_address"`
	ContractAddressSalt       string   `json:"contract_address_salt"`
	ClassHash                 Felt     `json:"class_hash"`
	ConstructorCalldata       []string `json:"constructor_calldata"`
	Signature                 []string `json:"signature"`
	ChainId                   Felt     `json:"chain_id"`
	FeeDataAvailabilityMode   Felt     `json:"fee_data_availability_mode"`
	NonceDataAvailabilityMode Felt     `json:"nonce_data_availability_mode"`
	PayMasterData             []Felt   `json:"paymaster_data"`
	Tip                       Felt     `json:"tip"`
}

func (d DeployAccount) toMap(version Felt) map[string]any {
	data := map[string]any{
		"max_fee":               d.MaxFee,
		"nonce":                 d.Nonce,
		"contract_address_salt": d.ContractAddressSalt,
		"signature":             d.Signature,
		"constructor_calldata":  d.ConstructorCalldata,
		"class_hash":            d.ClassHash,
		"contract_address":      d.ContractAddress,
	}

	switch version {
	case Version0, Version1, Version2:
		return data
	case Version3:
		data["chain_id"] = d.ChainId
		data["fee_data_availability_mode"] = d.FeeDataAvailabilityMode
		data["nonce_data_availability_mode"] = d.NonceDataAvailabilityMode
		data["tip"] = d.Tip
		data["paymaster_data"] = d.PayMasterData
	}

	return data
}

// L1Handler -
type L1Handler struct {
	Nonce              Felt     `json:"nonce"`
	ContractAddress    Felt     `json:"contract_address"`
	EntrypointSelector Felt     `json:"entry_point_selector"`
	Calldata           []string `json:"calldata"`
}

func (l1handler L1Handler) toMap() map[string]any {
	return map[string]any{
		"nonce":                l1handler.Nonce,
		"contract_address":     l1handler.ContractAddress,
		"entry_point_selector": l1handler.EntrypointSelector,
		"calldata":             l1handler.Calldata,
	}
}

// Transaction -
type Transaction struct {
	Type            string `json:"type"`
	Version         Felt   `json:"version"`
	TransactionHash Felt   `json:"transaction_hash,omitempty"`

	Body any `json:"-"`
}

// UnmarshalJSON -
func (t *Transaction) UnmarshalJSON(raw []byte) error {
	type buf Transaction
	if err := json.Unmarshal(raw, (*buf)(t)); err != nil {
		return err
	}

	switch t.Type {
	case TransactionTypeInvoke, TransactionTypeInvokeFunction:
		t.Body = &Invoke{}
	case TransactionTypeDeclare:
		t.Body = &Declare{}
	case TransactionTypeDeploy:
		t.Body = &Deploy{}
	case TransactionTypeDeployAccount:
		t.Body = &DeployAccount{}
	case TransactionTypeL1Handler:
		t.Body = &L1Handler{}
	default:
		return errors.Errorf("unknown transaction type: %s", t.Type)
	}

	return json.Unmarshal(raw, t.Body)
}

// MarshalJSON -
func (t *Transaction) MarshalJSON() ([]byte, error) {
	var m map[string]any

	switch t.Type {
	case TransactionTypeInvoke, TransactionTypeInvokeFunction:
		if invoke, ok := t.Body.(*Invoke); ok {
			m = invoke.toMap(t.Version)
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected InvokeV0 (non-pointer)")
		}
	case TransactionTypeDeclare:
		if declare, ok := t.Body.(*Declare); ok {
			m = declare.toMap(t.Version)
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected Declare (non-pointer)")
		}
	case TransactionTypeDeploy:
		if deploy, ok := t.Body.(*Deploy); ok {
			m = deploy.toMap()
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected Deploy (non-pointer)")
		}
	case TransactionTypeDeployAccount:
		if deploy, ok := t.Body.(*DeployAccount); ok {
			m = deploy.toMap(t.Version)
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected DeployAccount (non-pointer)")
		}
	case TransactionTypeL1Handler:
		if l1handler, ok := t.Body.(*L1Handler); ok {
			m = l1handler.toMap()
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected L1Handler (non-pointer)")
		}
	default:
		return nil, errors.Errorf("unknown transaction type: %s", t.Type)
	}

	m["type"] = t.Type
	m["version"] = t.Version

	return json.Marshal(m)
}
