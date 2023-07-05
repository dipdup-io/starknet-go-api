package abi

import (
	"fmt"

	"github.com/dipdup-net/indexer-sdk/pkg/jsonschema"
)

// JsonSchema -
type JsonSchema struct {
	Functions       map[string]JsonSchemaFunction    `json:"functions,omitempty"`
	L1Handlers      map[string]JsonSchemaFunction    `json:"l1_handlers,omitempty"`
	Constructors    map[string]JsonSchemaFunction    `json:"constructors,omitempty"`
	Events          map[string]JsonSchemaEvent       `json:"events,omitempty"`
	Structs         map[string]jsonschema.JSONSchema `json:"structs,omitempty"`
	Enums           map[string]jsonschema.JSONSchema `json:"enums,omitempty"`
	Interfaces      map[string]JsonSchemaInterface   `json:"interfaces,omitempty"`
	Implementations map[string]string                `json:"impls,omitempty"`
}

// JsonSchemaFunction -
type JsonSchemaFunction struct {
	Input  *jsonschema.JSONSchema `json:"input"`
	Output *jsonschema.JSONSchema `json:"output"`
}

func newJsonSchemaFunction(f FunctionItem, structs map[string]jsonschema.JSONSchema) JsonSchemaFunction {
	schema := JsonSchemaFunction{
		Input: &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
		},
		Output: &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
		},
	}

	buildJsonSchema(f.Inputs, structs, schema.Input.ObjectItem.Properties)
	for name := range schema.Input.ObjectItem.Properties {
		schema.Input.ObjectItem.Required = append(schema.Input.ObjectItem.Required, name)
	}
	buildJsonSchema(f.Outputs, structs, schema.Output.ObjectItem.Properties)
	for name := range schema.Output.ObjectItem.Properties {
		schema.Output.ObjectItem.Required = append(schema.Output.ObjectItem.Required, name)
	}
	return schema
}

// JsonSchemaEvent -
type JsonSchemaEvent struct {
	Keys   *jsonschema.JSONSchema `json:"keys,omitempty"`
	Data   *jsonschema.JSONSchema `json:"data,omitempty"`
	Inputs *jsonschema.JSONSchema `json:"inputs,omitempty"`
}

func newJsonSchemaEvent(f EventItem, structs map[string]jsonschema.JSONSchema) JsonSchemaEvent {
	schema := JsonSchemaEvent{}
	if f.Inputs != nil {
		schema.Inputs = &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
		}
		buildJsonSchema(f.Inputs, structs, schema.Inputs.ObjectItem.Properties)
		for name := range schema.Inputs.ObjectItem.Properties {
			schema.Inputs.ObjectItem.Required = append(schema.Inputs.ObjectItem.Required, name)
		}
	} else {
		schema.Keys = &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
		}
		schema.Data = &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
		}
		buildJsonSchema(f.Keys, structs, schema.Keys.ObjectItem.Properties)
		for name := range schema.Keys.ObjectItem.Properties {
			schema.Keys.ObjectItem.Required = append(schema.Keys.ObjectItem.Required, name)
		}
		buildJsonSchema(f.Data, structs, schema.Data.ObjectItem.Properties)
		for name := range schema.Data.ObjectItem.Properties {
			schema.Data.ObjectItem.Required = append(schema.Data.ObjectItem.Required, name)
		}
	}

	return schema
}

func newJsonSchemaForStruct(name string, item StructItem) jsonschema.JSONSchema {
	schema := jsonschema.JSONSchema{
		Title: name,
		ObjectItem: jsonschema.ObjectItem{
			Properties: make(map[string]jsonschema.JSONSchema),
			Required:   []string{},
		},
		InternalType: item.Type.Type,
		Type:         jsonschema.ItemTypeObject,
	}

	buildJsonSchemaForMembers(item.Members, map[string]jsonschema.JSONSchema{}, schema.ObjectItem.Properties)
	for name := range schema.ObjectItem.Properties {
		schema.ObjectItem.Required = append(schema.ObjectItem.Required, name)
	}

	return schema
}

func newJsonSchemaForEnum(name string, item EnumItem) jsonschema.JSONSchema {
	schema := jsonschema.JSONSchema{
		Title:        name,
		InternalType: item.Type.Type,
		Type:         jsonschema.ItemTypeString,
		Enum:         make([]any, 0),
	}

	for i := range item.Variants {
		schema.Enum = append(schema.Enum, item.Variants[i].Name)
	}

	return schema
}

// JsonSchemaInterface -
type JsonSchemaInterface map[string]JsonSchemaFunction

func newJsonSchemaForInterface(item InterfaceItem, structs map[string]jsonschema.JSONSchema) JsonSchemaInterface {
	schema := make(JsonSchemaInterface)

	for i := range item.Items {
		funcSchema := newJsonSchemaFunction(item.Items[i], structs)
		schema[item.Items[i].Name] = funcSchema
	}

	return schema
}

// JsonSchema -
func (abi Abi) JsonSchema() *JsonSchema {
	schema := &JsonSchema{
		Functions:       make(map[string]JsonSchemaFunction),
		L1Handlers:      make(map[string]JsonSchemaFunction),
		Constructors:    make(map[string]JsonSchemaFunction),
		Events:          make(map[string]JsonSchemaEvent),
		Structs:         make(map[string]jsonschema.JSONSchema),
		Enums:           make(map[string]jsonschema.JSONSchema),
		Interfaces:      make(map[string]JsonSchemaInterface),
		Implementations: make(map[string]string),
	}

	for name, typ := range abi.Structs {
		if typ == nil {
			continue
		}
		schema.Structs[name] = newJsonSchemaForStruct(name, *typ)
	}
	for name, typ := range abi.Enums {
		if typ == nil {
			continue
		}
		schema.Enums[name] = newJsonSchemaForEnum(name, *typ)
	}

	for name, typ := range abi.Functions {
		if typ == nil {
			continue
		}
		schema.Functions[name] = newJsonSchemaFunction(*typ, schema.Structs)
	}
	for name, typ := range abi.L1Handlers {
		if typ == nil {
			continue
		}
		schema.L1Handlers[name] = newJsonSchemaFunction(*typ, schema.Structs)
	}
	for name, typ := range abi.Constructor {
		if typ == nil {
			continue
		}
		schema.Constructors[name] = newJsonSchemaFunction(*typ, schema.Structs)
	}
	for name, typ := range abi.Events {
		if typ == nil {
			continue
		}
		schema.Events[name] = newJsonSchemaEvent(*typ, schema.Structs)
	}
	for name, typ := range abi.Interfaces {
		if typ == nil {
			continue
		}
		schema.Interfaces[name] = newJsonSchemaForInterface(*typ, schema.Structs)
	}
	for name, typ := range abi.Impls {
		if typ == nil {
			continue
		}
		schema.Implementations[name] = typ.InterfaceName
	}
	return schema
}

func buildJsonSchema(typ []Type, structs, out map[string]jsonschema.JSONSchema) {
	for i := range typ {
		name, js := getNameAndTypeForJsonSchema(typ[i], structs)
		out[name] = js
	}
}

func buildJsonSchemaForMembers(typ []Member, structs, out map[string]jsonschema.JSONSchema) {
	for i := range typ {
		name, js := getNameAndTypeForJsonSchema(typ[i].Type, structs)
		out[name] = js
	}
}

func getNameAndTypeForJsonSchema(typ Type, structs map[string]jsonschema.JSONSchema) (string, jsonschema.JSONSchema) {
	switch typ.Type {
	case coreTypeFelt, coreTypeFelt252, coreTypeContractAddress, coreTypeClassHash:
		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeString,
			Title:        typ.Name,
			InternalType: typ.Type,
		}
	case coreTypeU8, coreTypeU16, coreTypeU32, coreTypeU64, coreTypeU128, coreTypeU256:
		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeInteger,
			Title:        typ.Name,
			InternalType: typ.Type,
		}
	case coreTypeBool:
		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeBoolean,
			Title:        typ.Name,
			InternalType: typ.Type,
		}
	case coreTypeOption:
		optionType := unwrapOptionType(typ.Type)
		_, item := getNameAndTypeForJsonSchema(Type{
			Name: fmt.Sprintf("%s_item", typ.Name),
			Type: optionType,
		}, structs)

		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeObject,
			Title:        typ.Name,
			InternalType: typ.Type,
			OneOf: []*jsonschema.JSONSchema{
				{
					Type: jsonschema.ItemTypeNull,
				},
				&item,
			},
		}
	case coreTypeECPoint:
		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeObject,
			Title:        typ.Name,
			InternalType: typ.Type,
			ObjectItem: jsonschema.ObjectItem{
				Properties: map[string]jsonschema.JSONSchema{
					"x": {
						Type:         jsonschema.ItemTypeString,
						Title:        "x",
						InternalType: coreTypeFelt,
					},
					"y": {
						Type:         jsonschema.ItemTypeString,
						Title:        "y",
						InternalType: coreTypeFelt,
					},
				},
			},
		}
	default:
		if isTypeArray(typ.Type) {
			itemType := unwrapArrayType(typ.Type)
			_, item := getNameAndTypeForJsonSchema(Type{
				Name: fmt.Sprintf("%s_item", typ.Name),
				Type: itemType,
			}, structs)

			return typ.Name, jsonschema.JSONSchema{
				Type: jsonschema.ItemTypeArray,
				ArrayItem: jsonschema.ArrayItem{
					Items: []jsonschema.JSONSchema{
						item,
					},
				},
			}
		}

		if t, ok := structs[typ.Type]; ok {
			return typ.Name, t
		} else {
			return typ.Name, jsonschema.JSONSchema{
				Type:         jsonschema.ItemTypeString,
				Title:        typ.Name,
				InternalType: typ.Type,
			}
		}
	}
}
