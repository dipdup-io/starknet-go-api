package abi

import (
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"

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
		case InterfaceType:
			items = append(items, &InterfaceItem{})

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
	a.Impls = make(map[string]*ImplItem)
	a.Interfaces = make(map[string]*InterfaceItem)

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
		case ImplType:
			item := items[i].(*ImplItem)
			a.Impls[types[i].Name] = item
		case InterfaceType:
			item := items[i].(*InterfaceItem)
			a.Interfaces[types[i].Name] = item
			for i := range item.Items {
				selector := encoding.GetSelectorFromName(item.Items[i].Name)
				switch item.Items[i].Type.Type {
				case FunctionType:
					a.Functions[item.Items[i].Name] = &item.Items[i]
					a.FunctionsBySelector[selector] = &item.Items[i]
				default:
					log.Warn().
						Str("typ", item.Items[i].Type.Type).
						Msgf("unknown interface item type: %s %s", item.Items[i].Type.Name, item.Items[i].Type.Type)
				}
			}
		default:
			return errors.Errorf("unknown abi type: %s", types[i].Type)
		}
	}

	return nil
}
