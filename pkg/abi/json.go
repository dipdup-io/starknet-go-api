package abi

import (
	"github.com/goccy/go-json"

	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
	"github.com/pkg/errors"
)

// UnmarshalJSON -
func (a *Abi) UnmarshalJSON(raw []byte) error {
	var types []Type
	if err := json.Unmarshal(raw, &types); err != nil {
		return err
	}

	items := make([]any, 0)
	for i := range types {
		switch types[i].Type {
		case ConstructorType, FunctionType, L1HandlerType:
			items = append(items, &FunctionItem{})
		case EventType:
			items = append(items, &EventItem{})
		case StructType:
			items = append(items, &StructItem{})
		case EnumType:
			items = append(items, &EnumItem{})
		case ImplType:
			items = append(items, &ImplItem{})

		default:
			return errors.Errorf("unknown abi type: %s", types[i].Type)
		}
	}

	if err := json.Unmarshal(raw, &items); err != nil {
		return err
	}

	a.Constructor = make(map[string]*FunctionItem)
	a.Events = make(map[string]*EventItem)
	a.Functions = make(map[string]*FunctionItem)
	a.L1Handlers = make(map[string]*FunctionItem)
	a.Structs = make(map[string]*StructItem)
	a.Enums = make(map[string]*EnumItem)

	a.ConstructorBySelector = make(map[string]*FunctionItem)
	a.EventsBySelector = make(map[string]*EventItem)
	a.FunctionsBySelector = make(map[string]*FunctionItem)
	a.L1HandlersBySelector = make(map[string]*FunctionItem)
	a.StructsBySelector = make(map[string]*StructItem)

	for i := range types {
		selector := encoding.GetSelectorFromName(types[i].Name)
		switch types[i].Type {
		case L1HandlerType:
			item := items[i].(*FunctionItem)
			a.L1Handlers[types[i].Name] = item
			a.L1HandlersBySelector[selector] = item
		case FunctionType:
			item := items[i].(*FunctionItem)
			a.Functions[types[i].Name] = item
			a.FunctionsBySelector[selector] = item
		case ConstructorType:
			item := items[i].(*FunctionItem)
			a.Constructor[types[i].Name] = item
			a.ConstructorBySelector[selector] = item
		case EventType:
			item := items[i].(*EventItem)
			a.Events[types[i].Name] = item
			a.EventsBySelector[selector] = item
		case StructType:
			item := items[i].(*StructItem)
			a.Structs[types[i].Name] = item
			a.StructsBySelector[selector] = item
		case EnumType:
			item := items[i].(*EnumItem)
			a.Enums[types[i].Name] = item
		default:
			return errors.Errorf("unknown abi type: %s", types[i].Type)
		}
	}

	return nil
}
