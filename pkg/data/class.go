package data

import (
	stdJSON "encoding/json"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Class -
type Class struct {
	EntryPointsByType EntrypointsByType  `json:"entry_points_by_type"`
	Abi               Abi                `json:"-"`
	RawAbi            stdJSON.RawMessage `json:"abi"`
}

// GetAbi -
func (c *Class) GetAbi() (Abi, error) {
	var abi Abi
	err := json.Unmarshal(c.RawAbi, &abi)
	return abi, err
}

// Handler -
type Handler struct {
	Offset   string `json:"offset"`
	Selector string `json:"selector"`
}

// Type -
type Type struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// EntrypointsByType -
type EntrypointsByType struct {
	CONSTRUCTOR []Handler `json:"CONSTRUCTOR"`
	EXTERNAL    []Handler `json:"EXTERNAL"`
	L1HANDLER   []Handler `json:"L1_HANDLER"`
}

// FunctionAbiItem -
type FunctionAbiItem struct {
	Type

	Inputs  []Type `json:"inputs"`
	Outputs []Type `json:"outputs"`
}

// FunctionAbiItem -
type EventAbiItem struct {
	Type

	Data []Type `json:"data"`
	Keys []Type `json:"keys"`
}

// StructAbiItem -
type StructAbiItem struct {
	Type

	Size    uint64   `json:"size"`
	Members []Member `json:"members"`
}

// Member -
type Member struct {
	Type

	Offset uint64 `json:"offset"`
}

// Abi -
type Abi struct {
	Functions   map[string]*FunctionAbiItem `json:"-"`
	L1Handlers  map[string]*FunctionAbiItem `json:"-"`
	Constructor map[string]*FunctionAbiItem `json:"-"`
	Events      map[string]*EventAbiItem    `json:"-"`
	Structs     map[string]*StructAbiItem   `json:"-"`
}

// UnmarshalJSON -
func (a *Abi) UnmarshalJSON(raw []byte) error {
	var types []Type
	if err := json.Unmarshal(raw, &types); err != nil {
		return err
	}

	items := make([]any, 0)
	for i := range types {
		switch types[i].Type {
		case AbiConstructorType, AbiFunctionType, AbiL1HandlerType:
			items = append(items, &FunctionAbiItem{})
		case AbiEventType:
			items = append(items, &EventAbiItem{})
		case AbiStructType:
			items = append(items, &StructAbiItem{})
		default:
			return errors.Errorf("unknown abi type: %s", types[i].Type)
		}
	}

	if err := json.Unmarshal(raw, &items); err != nil {
		return err
	}

	a.Constructor = make(map[string]*FunctionAbiItem)
	a.Events = make(map[string]*EventAbiItem)
	a.Functions = make(map[string]*FunctionAbiItem)
	a.L1Handlers = make(map[string]*FunctionAbiItem)
	a.Structs = make(map[string]*StructAbiItem)

	for i := range types {
		switch types[i].Type {
		case AbiL1HandlerType:
			a.L1Handlers[types[i].Name] = items[i].(*FunctionAbiItem)
		case AbiFunctionType:
			a.Functions[types[i].Name] = items[i].(*FunctionAbiItem)
		case AbiConstructorType:
			a.Constructor[types[i].Name] = items[i].(*FunctionAbiItem)
		case AbiEventType:
			a.Events[types[i].Name] = items[i].(*EventAbiItem)
		case AbiStructType:
			a.Structs[types[i].Name] = items[i].(*StructAbiItem)
		default:
			return errors.Errorf("unknown abi type: %s", types[i].Type)
		}
	}

	return nil
}
