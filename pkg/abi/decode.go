package abi

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
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
	return strings.HasSuffix(typ, "*") ||
		strings.HasPrefix(typ, coreTypeArray) ||
		strings.HasPrefix(typ, coreTypeSpan)
}

func isTypeTuple(typ string) bool {
	l := len(typ)
	if l < 2 {
		return false
	}
	return typ[0] == '(' && typ[l-1] == ')'
}

func isTypeOption(typ string) bool {
	return strings.HasPrefix(typ, coreTypeOption)
}

func unwrapArrayType(typ string) string {
	switch {
	case strings.HasSuffix(typ, "*"):
		return strings.TrimSuffix(typ, "*")
	case strings.HasPrefix(typ, coreTypeArray):
		s := strings.TrimPrefix(typ, coreTypeArray+"::<")
		return strings.TrimSuffix(s, ">")
	case strings.HasPrefix(typ, coreTypeSpan):
		s := strings.TrimPrefix(typ, coreTypeSpan+"::<")
		return strings.TrimSuffix(s, ">")
	}
	return typ
}

func unwrapOptionType(typ string) string {
	s := strings.TrimPrefix(typ, coreTypeOption+"::<")
	return strings.TrimSuffix(s, ">")
}

// DecodeExecuteCallData -
func DecodeExecuteCallData(calldata []string) (map[string]any, error) {
	return DecodeFunctionCallData(calldata, ExecuteFunction, map[string]*StructItem{
		"CallArray": &CallArray,
	}, nil)
}

// DecodeExecuteResult -
func DecodeExecuteResult(result []string) (map[string]any, error) {
	return DecodeFunctionCallData(result, ExecuteFunction, map[string]*StructItem{
		"CallArray": &CallArray,
	}, nil)
}

// DecodeChangeModulesCallData -
func DecodeChangeModulesCallData(calldata []string) (map[string]any, error) {
	return DecodeFunctionCallData(calldata, ChangeModules, map[string]*StructItem{
		"ModuleFunctionAction": &ModuleFunctionAction,
	}, nil)
}

// DecodeChangeModulesResult -
func DecodeChangeModulesResult(result []string) (map[string]any, error) {
	return DecodeFunctionCallData(result, ChangeModules, map[string]*StructItem{
		"ModuleFunctionAction": &ModuleFunctionAction,
	}, nil)
}

// DecodeFunctionCallData -
func DecodeFunctionCallData(calldata []string, typ FunctionItem, structs map[string]*StructItem, enums map[string]*EnumItem) (map[string]any, error) {

	var (
		result = make(map[string]any, 0)
		tail   = calldata
		err    error
	)

	for _, input := range typ.Inputs {
		tail, err = decodeItem(tail, input, structs, enums, result)
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

// DecodeEventData -
func DecodeEventData(data []string, typ EventItem, structs map[string]*StructItem, enums map[string]*EnumItem) (map[string]any, error) {
	var (
		result = make(map[string]any, 0)
		tail   = data
		err    error
	)

	for _, input := range typ.Data {
		tail, err = decodeItem(tail, input, structs, enums, result)
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

// DecodeFunctionResult -
func DecodeFunctionResult(data []string, typ FunctionItem, structs map[string]*StructItem, enums map[string]*EnumItem) (map[string]any, error) {
	var (
		result = make(map[string]any, 0)
		tail   = data
		err    error
	)

	for _, output := range typ.Outputs {
		tail, err = decodeItem(tail, output, structs, enums, result)
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

func decodeItem(calldata []string, input Type, structs map[string]*StructItem, enums map[string]*EnumItem, result map[string]any) ([]string, error) {
	enum, hasEnum := enums[input.Type]
	str, hasStruct := structs[input.Type]
	switch {
	case isLenField(input.Name):
		if len(calldata) == 0 {
			return nil, ErrTooShortCallData
		}
		result[input.Name] = calldata[0]
		return calldata[1:], nil

	case hasEnum:
		if len(calldata) == 0 {
			return nil, ErrTooShortCallData
		}
		enumIdx, err := strconv.ParseInt(calldata[0], 0, 64)
		if err != nil {
			return nil, err
		}
		if int(enumIdx) > len(enum.Variants)-1 {
			return nil, errors.Errorf("too big enum index: %d", enumIdx)
		}
		variant := enum.Variants[enumIdx]
		obj := make(map[string]any)
		tail, err := decodeItem(calldata[1:], variant, structs, enums, obj)
		if err != nil {
			return nil, err
		}
		result[input.Name] = obj
		return tail, nil

	case hasStruct:
		obj := make(map[string]any)
		tail := calldata
		var err error
		for i := range str.Members {
			tail, err = decodeItem(tail, str.Members[i].Type, structs, enums, obj)
			if err != nil {
				return nil, err
			}
		}
		result[input.Name] = obj
		return tail, nil

	case isTypeArray(input.Type):
		var iLength int
		switch {
		case strings.HasSuffix(input.Type, "*"):
			lengthHex, ok := result[fmt.Sprintf("%s_len", input.Name)]
			if !ok {
				return nil, errors.Wrap(ErrNoLenField, input.Name)
			}
			length, err := strconv.ParseInt(lengthHex.(string), 0, 64)
			if err != nil {
				return nil, errors.Wrap(err, input.Name)
			}
			iLength = int(length)
		case strings.HasPrefix(input.Type, coreTypeArray) || strings.HasPrefix(input.Type, coreTypeSpan):
			if len(calldata) == 0 {
				return nil, ErrTooShortCallData
			}
			length, err := strconv.ParseInt(calldata[0], 0, 64)
			if err != nil {
				return nil, errors.Wrap(err, input.Name)
			}
			iLength = int(length)
			calldata = calldata[1:]
		}

		if iLength == 0 {
			result[input.Name] = []any{}
			return calldata, nil
		}

		if len(calldata) < iLength {
			return nil, ErrTooShortCallData
		}

		parsed := make([]any, iLength)
		tail := calldata
		var err error

		for i := 0; i < iLength; i++ {
			obj := make(map[string]any)
			tail, err = decodeItem(tail, Type{
				Name: fmt.Sprintf("array_item_%d", i),
				Type: unwrapArrayType(input.Type),
			}, structs, enums, obj)
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
			tail, err = decodeItem(tail, tupleItems[i].Type, structs, enums, obj)
			if err != nil {
				return nil, err
			}
		}
		result[input.Name] = obj
		return tail, nil

	case input.Type == coreTypeU256:
		if len(calldata) < 2 {
			return nil, ErrTooShortCallData
		}
		bigInt, err := DecodeUint256(calldata[0], calldata[1])
		if err != nil {
			result[input.Name] = []string{
				calldata[0], calldata[1],
			}
		} else {
			result[input.Name] = bigInt.String()
		}
		return calldata[2:], nil

	case isTypeOption(input.Type):
		if len(calldata) < 1 {
			return nil, ErrTooShortCallData
		}

		switch calldata[0] {
		case optionNone:
			result[input.Name] = nil
			return calldata[1:], nil
		case optionSome:
			if len(calldata) < 2 {
				return nil, ErrTooShortCallData
			}
			obj := make(map[string]any)
			tail, err := decodeItem(calldata[1:], Type{
				Name: fmt.Sprintf("%s_some", input.Name),
				Type: unwrapOptionType(input.Type),
			}, structs, enums, obj)
			if err != nil {
				return nil, err
			}
			result[input.Name] = obj
			return tail, nil
		}

		return calldata, nil

	case input.Type == coreTypeECPoint:
		if len(calldata) < 2 {
			return nil, ErrTooShortCallData
		}
		result[input.Name] = map[string]string{
			"x": calldata[0],
			"y": calldata[1],
		}
		return calldata[2:], nil

	default:
		if len(calldata) == 0 {
			return nil, ErrTooShortCallData
		}
		decodeSimpleType(input, calldata[0], result)
		return calldata[1:], nil
	}
}

func decodeSimpleType(input Type, value string, result map[string]any) {
	switch input.Type {
	case coreTypeBool:
		b, err := strconv.ParseBool(encoding.TrimHex(value))
		if err != nil {
			result[input.Name] = value
			return
		}
		result[input.Name] = b
	case coreTypeU8, coreTypeU16, coreTypeU32, coreTypeU64:
		u, err := strconv.ParseUint(encoding.TrimHex(value), 16, 64)
		if err != nil {
			result[input.Name] = value
			return
		}
		result[input.Name] = u
	case coreTypeU128:
		bigInt, err := decodeUint128(value)
		if err != nil {
			result[input.Name] = value
			return
		}
		result[input.Name] = bigInt.String()
	default:
		result[input.Name] = value
	}
}

// DecodeUint256 - parser 2 words and returns big.Int
func DecodeUint256(low, high string) (*big.Int, error) {
	highInt, err := decodeUint128(high)
	if err != nil {
		return nil, errors.Wrap(err, "invalid high of uint256")
	}
	lowInt, err := decodeUint128(low)
	if err != nil {
		return nil, errors.Wrap(err, "invalid low of uint256")
	}
	return highInt.Lsh(highInt, 128).Add(highInt, lowInt), nil
}

func decodeUint128(value string) (*big.Int, error) {
	bigInt, ok := big.NewInt(0).SetString(value, 0)
	if !ok {
		return nil, errors.Errorf("invalid uint128: %s", value)
	}
	return bigInt, nil
}
