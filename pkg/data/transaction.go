package data

import (
	"github.com/pkg/errors"
)

// InvokeV0 -
type InvokeV0 struct {
	MaxFee             string   `json:"max_fee"`
	Nonce              string   `json:"nonce"`
	ContractAddress    string   `json:"contract_address"`
	EntrypointSelector string   `json:"entry_point_selector"`
	Signature          []string `json:"signature"`
	Calldata           []string `json:"calldata"`
}

func (i InvokeV0) toMap() map[string]any {
	return map[string]any{
		"max_fee":              i.MaxFee,
		"nonce":                i.Nonce,
		"contract_address":     i.ContractAddress,
		"entry_point_selector": i.EntrypointSelector,
		"signature":            i.Signature,
		"calldata":             i.Calldata,
	}
}

// InvokeV1 -
type InvokeV1 struct {
	MaxFee        string   `json:"max_fee"`
	Nonce         string   `json:"nonce"`
	SenderAddress string   `json:"sender_address"`
	Signature     []string `json:"signature"`
	Calldata      []string `json:"calldata"`
}

func (i InvokeV1) toMap() map[string]any {
	return map[string]any{
		"max_fee":        i.MaxFee,
		"nonce":          i.Nonce,
		"sender_address": i.SenderAddress,
		"signature":      i.Signature,
		"calldata":       i.Calldata,
	}
}

// Declare -
type Declare struct {
	MaxFee        string   `json:"max_fee"`
	Nonce         string   `json:"nonce"`
	SenderAddress string   `json:"sender_address"`
	Signature     []string `json:"signature"`
	ContractClass Class    `json:"contract_class,omitempty"`
	ClassHash     string   `json:"class_hash,omitempty"`
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
	ClassHash           string   `json:"class_hash,omitempty"`
	ContractClass       Class    `json:"contract_class,omitempty"`
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
	MaxFee              string   `json:"max_fee"`
	Nonce               string   `json:"nonce"`
	ContractAddressSalt string   `json:"contract_address_salt"`
	ClassHash           string   `json:"class_hash"`
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
	}
}

// L1Handler -
type L1Handler struct {
	Nonce              string   `json:"nonce"`
	ContractAddress    string   `json:"contract_address"`
	EntrypointSelector string   `json:"entry_point_selector"`
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
	Version         string `json:"version"`
	TransactionHash string `json:"transaction_hash,omitempty"`

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

		switch t.Version {
		case Version0:
			t.Body = &InvokeV0{}
		case Version1:
			t.Body = &InvokeV1{}
		default:
			return errors.Errorf("unknown transaction version: %s", t.Version)
		}

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

		switch t.Version {
		case Version0:
			if invoke, ok := t.Body.(InvokeV0); ok {
				m = invoke.toMap()
			} else {
				return nil, errors.Errorf("invalid invoke transaction type: expected InvokeV0 (non-pointer)")
			}
		case Version1:
			if invoke, ok := t.Body.(InvokeV1); ok {
				m = invoke.toMap()
			} else {
				return nil, errors.Errorf("invalid invoke transaction type: expected InvokeV1 (non-pointer)")
			}
		default:
			return nil, errors.Errorf("unknown transaction version: %s", t.Version)
		}

	case TransactionTypeDeclare:
		if declare, ok := t.Body.(Declare); ok {
			m = declare.toMap()
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected Declare (non-pointer)")
		}
	case TransactionTypeDeploy:
		if deploy, ok := t.Body.(Deploy); ok {
			m = deploy.toMap()
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected Deploy (non-pointer)")
		}
	case TransactionTypeDeployAccount:
		if deploy, ok := t.Body.(DeployAccount); ok {
			m = deploy.toMap()
		} else {
			return nil, errors.Errorf("invalid invoke transaction type: expected DeployAccount (non-pointer)")
		}
	case TransactionTypeL1Handler:
		if l1handler, ok := t.Body.(L1Handler); ok {
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
