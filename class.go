package starknetgoapi

import "github.com/pkg/errors"

// Class -
type Class struct {
	Program           string            `json:"program"`
	EntryPointsByType EntrypointsByType `json:"entry_points_by_type"`
	Abi               Abi               `json:"abi"`
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
	Functions   []*FunctionAbiItem `json:"-"`
	L1Handlers  []*FunctionAbiItem `json:"-"`
	Constructor []*FunctionAbiItem `json:"-"`
	Events      []*EventAbiItem    `json:"-"`
	Structs     []*StructAbiItem   `json:"-"`
}

// UnmarshalJSON -
func (a *Abi) UnmarshalJSON(data []byte) error {
	var types []Type
	if err := json.Unmarshal(data, &types); err != nil {
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

	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	a.Constructor = make([]*FunctionAbiItem, 0)
	a.Events = make([]*EventAbiItem, 0)
	a.Functions = make([]*FunctionAbiItem, 0)
	a.L1Handlers = make([]*FunctionAbiItem, 0)
	a.Structs = make([]*StructAbiItem, 0)

	for i := range types {
		switch types[i].Type {
		case AbiL1HandlerType:
			a.L1Handlers = append(a.L1Handlers, items[i].(*FunctionAbiItem))
		case AbiFunctionType:
			a.Functions = append(a.Functions, items[i].(*FunctionAbiItem))
		case AbiConstructorType:
			a.Constructor = append(a.Constructor, items[i].(*FunctionAbiItem))
		case AbiEventType:
			a.Events = append(a.Events, items[i].(*EventAbiItem))
		case AbiStructType:
			a.Structs = append(a.Structs, items[i].(*StructAbiItem))
		default:
			return errors.Errorf("unknown abi type: %s", types[i].Type)
		}
	}

	return nil
}
