package abi

import "github.com/goccy/go-json"

// abi types
const (
	FunctionType    = "function"
	L1HandlerType   = "l1_handler"
	ConstructorType = "constructor"
	EventType       = "event"
	StructType      = "struct"
	EnumType        = "enum"
	ImplType        = "impl"
	InterfaceType   = "interface"
)

// core types
const (
	coreTypeBool            = "core::bool"
	coreTypeU8              = "core::integer::u8"
	coreTypeU16             = "core::integer::u16"
	coreTypeU32             = "core::integer::u32"
	coreTypeU64             = "core::integer::u64"
	coreTypeU128            = "core::integer::u128"
	coreTypeU256            = "core::integer::u256"
	coreTypeFelt            = "felt"
	coreTypeFelt252         = "core::felt252"
	coreTypeContractAddress = "core::starknet::contract_address::ContractAddress"
	coreTypeArray           = "core::array::Array"
	coreTypeSpan            = "core::array::Span"
	coreTypeClassHash       = "core::starknet::class_hash::ClassHash"
	coreTypeOption          = "core::option::Option"
	coreTypeECPoint         = "core::ec::EcPoint"
)

const (
	optionSome = "0x0"
	optionNone = "0x1"
)

// Abi -
type Abi struct {
	Functions   map[string]*FunctionItem  `json:"-"`
	L1Handlers  map[string]*FunctionItem  `json:"-"`
	Constructor map[string]*FunctionItem  `json:"-"`
	Events      map[string]*EventItem     `json:"-"`
	Structs     map[string]*StructItem    `json:"-"`
	Enums       map[string]*EnumItem      `json:"-"`
	Impls       map[string]*ImplItem      `json:"-"`
	Interfaces  map[string]*InterfaceItem `json:"-"`

	FunctionsBySelector   map[string]*FunctionItem `json:"-"`
	L1HandlersBySelector  map[string]*FunctionItem `json:"-"`
	ConstructorBySelector map[string]*FunctionItem `json:"-"`
	EventsBySelector      map[string]*EventItem    `json:"-"`
	StructsBySelector     map[string]*StructItem   `json:"-"`

	Names map[string]string `json:"-"`
}

// Type -
type Type struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Kind string `json:"kind,omitempty"`
}

// FunctionItem -
type FunctionItem struct {
	Type

	Inputs  []Type `json:"inputs"`
	Outputs []Type `json:"outputs"`
}

// EventItem -
type EventItem struct {
	Type

	Members []Type `json:"members,omitempty"`
	Data    []Type `json:"data"`
	Keys    []Type `json:"keys"`
	Inputs  []Type `json:"inputs"`
}

type Members struct {
}

func (item *EventItem) UnmarshalJSON(data []byte) error {
	type buf EventItem
	if err := json.Unmarshal(data, (*buf)(item)); err != nil {
		return err
	}
	if item.Kind == StructType {
		item.Data = make([]Type, 0)
		item.Keys = make([]Type, 0)
		for i := range item.Members {
			switch item.Members[i].Kind {
			case "data":
				item.Data = append(item.Data, item.Members[i])
			case "keys":
				item.Keys = append(item.Keys, item.Members[i])
			}
		}
	}
	return nil
}

// StructItem -
type StructItem struct {
	Type

	Size    uint64   `json:"size"`
	Members []Member `json:"members"`
}

// EnumItem -
type EnumItem struct {
	Type

	Variants []Type `json:"variants"`
}

// ImplItem -
type ImplItem struct {
	Type

	InterfaceName string `json:"interface_name"`
}

// InterfaceItem -
type InterfaceItem struct {
	Type

	Items []FunctionItem `json:"items"`
}

// Member -
type Member struct {
	Type

	Offset uint64 `json:"offset"`
}

// entrypoint types
const (
	EntrypointTypeExternal    = "EXTERNAL"
	EntrypointTypeConstructor = "CONSTRUCTOR"
	EntrypointTypeL1Handler   = "L1_HANDLER"
)
