package data

import (
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

// Invoke -
type Invoke struct {
	MaxFee             Felt     `json:"max_fee"`
	Nonce              Felt     `json:"nonce"`
	ContractAddress    Felt     `json:"contract_address"`
	EntrypointSelector Felt     `json:"entry_point_selector"`
	SenderAddress      Felt     `json:"sender_address"`
	Signature          []string `json:"signature"`
	Calldata           []string `json:"calldata"`
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
	}

	return data
}

// Declare -
type Declare struct {
	MaxFee          Felt     `json:"max_fee"`
	Nonce           Felt     `json:"nonce"`
	SenderAddress   Felt     `json:"sender_address"`
	ContractAddress Felt     `json:"contract_address"`
	Signature       []string `json:"signature"`
	ContractClass   Class    `json:"contract_class,omitempty"`
	ClassHash       Felt     `json:"class_hash,omitempty"`
}

func (d Declare) toMap() map[string]any {
	return map[string]any{
		"max_fee":        d.MaxFee,
		"nonce":          d.Nonce,
		"sender_address": d.SenderAddress,
		"signature":      d.Signature,
		"contract_class": d.ContractClass,
	}
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
	MaxFee              Felt     `json:"max_fee"`
	Nonce               Felt     `json:"nonce"`
	ContractAddress     Felt     `json:"contract_address"`
	ContractAddressSalt string   `json:"contract_address_salt"`
	ClassHash           Felt     `json:"class_hash"`
	ConstructorCalldata []string `json:"constructor_calldata"`
	Signature           []string `json:"signature"`
}

func (d DeployAccount) toMap() map[string]any {
	return map[string]any{
		"max_fee":               d.MaxFee,
		"nonce":                 d.Nonce,
		"contract_address_salt": d.ContractAddressSalt,
		"signature":             d.Signature,
		"constructor_calldata":  d.ConstructorCalldata,
		"class_hash":            d.ClassHash,
		"contract_address":      d.ContractAddress,
	}
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
			m = declare.toMap()
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
			m = deploy.toMap()
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
