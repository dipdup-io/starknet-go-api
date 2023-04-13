package abi

import (
	"fmt"
	"strings"

	"github.com/dipdup-net/indexer-sdk/pkg/jsonschema"
)

// JsonSchema -
type JsonSchema struct {
	Functions    map[string]JsonSchemaFunction    `json:"functions"`
	L1Handlers   map[string]JsonSchemaFunction    `json:"l1_handlers"`
	Constructors map[string]JsonSchemaFunction    `json:"constructors"`
	Events       map[string]JsonSchemaEvent       `json:"events"`
	Structs      map[string]jsonschema.JSONSchema `json:"structs"`
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
	Keys *jsonschema.JSONSchema `json:"keys"`
	Data *jsonschema.JSONSchema `json:"data"`
}

func newJsonSchemaEvent(f EventItem, structs map[string]jsonschema.JSONSchema) JsonSchemaEvent {
	schema := JsonSchemaEvent{
		Keys: &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
		},
		Data: &jsonschema.JSONSchema{
			Type: jsonschema.ItemTypeObject,
			ObjectItem: jsonschema.ObjectItem{
				Properties: make(map[string]jsonschema.JSONSchema),
				Required:   []string{},
			},
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

// JsonSchema -
func (abi Abi) JsonSchema() *JsonSchema {
	schema := &JsonSchema{
		Functions:    make(map[string]JsonSchemaFunction),
		L1Handlers:   make(map[string]JsonSchemaFunction),
		Constructors: make(map[string]JsonSchemaFunction),
		Events:       make(map[string]JsonSchemaEvent),
		Structs:      make(map[string]jsonschema.JSONSchema),
	}

	for name, typ := range abi.Structs {
		if typ == nil {
			continue
		}
		schema.Structs[name] = newJsonSchemaForStruct(name, *typ)
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
	case "felt", "core::felt252", "core::starknet::contract_address::ContractAddress":
		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeString,
			Title:        typ.Name,
			InternalType: typ.Type,
		}
	case "core::integer::u8", "core::integer::u128", "core::integer::u256":
		return typ.Name, jsonschema.JSONSchema{
			Type:         jsonschema.ItemTypeInteger,
			Title:        typ.Name,
			InternalType: typ.Type,
		}
	default:
		if strings.HasPrefix(typ.Type, "*") {
			itemType := strings.TrimPrefix(typ.Type, "*")
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
				InternalType: typ.Type,
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
