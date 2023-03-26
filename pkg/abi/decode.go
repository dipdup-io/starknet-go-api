package abi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// errors
var (
	ErrTooShortCallData = errors.New("calldata is too short")
	ErrInvalidTupleType = errors.New("invalid tuple type")
	ErrNoLenField       = errors.New("can't find array length field")
)

func isLenField(name string) bool {
	return strings.HasSuffix(name, "_len")
}

func isTypeArray(typ string) bool {
	return strings.Contains(typ, "*")
}

// func isTypeFelt(typ string) bool {
// 	return typ == "felt"
// }

// func isTypeFeltArray(typ string) bool {
// 	return typ == "felt*"
// }

func isTypeTuple(typ string) bool {
	l := len(typ)
	if l < 2 {
		return false
	}
	return typ[0] == '(' && typ[l-1] == ')'
}

// DecodeExecuteCallData -
func DecodeExecuteCallData(calldata []string) (map[string]any, error) {
	return DecodeFunctionCallData(calldata, ExecuteFunction, map[string]*StructItem{
		"CallArray": &CallArray,
	})
}

// DecodeChangeModulesCallData -
func DecodeChangeModulesCallData(calldata []string) (map[string]any, error) {
	return DecodeFunctionCallData(calldata, ChangeModules, map[string]*StructItem{
		"ModuleFunctionAction": &ModuleFunctionAction,
	})
}

// DecodeFunctionCallData -
func DecodeFunctionCallData(calldata []string, typ FunctionItem, structs map[string]*StructItem) (map[string]any, error) {

	var (
		result = make(map[string]any, 0)
		tail   = calldata
		err    error
	)

	for _, input := range typ.Inputs {
		tail, err = decodeItem(tail, input, structs, result)
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

// DecodeEventData -
func DecodeEventData(data []string, typ EventItem, structs map[string]*StructItem) (map[string]any, error) {
	var (
		result = make(map[string]any, 0)
		tail   = data
		err    error
	)

	for _, input := range typ.Data {
		tail, err = decodeItem(tail, input, structs, result)
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

func decodeItem(calldata []string, input Type, structs map[string]*StructItem, result map[string]any) ([]string, error) {
	str, hasStruct := structs[input.Type]
	switch {
	case isLenField(input.Name):
		if len(calldata) == 0 {
			return nil, ErrTooShortCallData
		}
		result[input.Name] = calldata[0]
		return calldata[1:], nil

	case hasStruct:
		obj := make(map[string]any)
		tail := calldata
		var err error
		for i := range str.Members {
			tail, err = decodeItem(tail, str.Members[i].Type, structs, obj)
			if err != nil {
				return nil, err
			}
		}
		result[input.Name] = obj
		return tail, nil

	case isTypeArray(input.Type):
		lengthHex, ok := result[fmt.Sprintf("%s_len", input.Name)]
		if !ok {
			return nil, errors.Wrap(ErrNoLenField, input.Name)
		}
		length, err := strconv.ParseInt(lengthHex.(string), 0, 64)
		if err != nil {
			return nil, errors.Wrap(err, input.Name)
		}
		iLength := int(length)

		if iLength == 0 {
			return calldata, nil
		}

		if len(calldata) < iLength {
			return nil, ErrTooShortCallData
		}

		parsed := make([]any, iLength)
		tail := calldata
		for i := 0; i < iLength; i++ {
			obj := make(map[string]any)
			tail, err = decodeItem(tail, Type{
				Name: fmt.Sprintf("array_item_%d", i),
				Type: strings.TrimSuffix(input.Type, "*"),
			}, structs, obj)
			if err != nil {
				return nil, err
			}
			for _, value := range obj {
				parsed[i] = value
			}
		}
		result[input.Name] = parsed

		return tail, nil

	case isTypeTuple(input.Type):
		tupleItems, err := extractTupleTypes(input.Type)
		if err != nil {
			return nil, err
		}

		obj := make(map[string]any)
		tail := calldata
		for i := range tupleItems {
			tail, err = decodeItem(tail, tupleItems[i].Type, structs, obj)
			if err != nil {
				return nil, err
			}
		}
		result[input.Name] = obj
		return tail, nil

	default:
		if len(calldata) == 0 {
			return nil, ErrTooShortCallData
		}
		result[input.Name] = calldata[0]
		return calldata[1:], nil
	}
}
