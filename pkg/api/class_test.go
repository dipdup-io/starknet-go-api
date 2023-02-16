package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbi_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Abi
		wantErr bool
	}{
		{
			name: "test 1",
			data: []byte(`[
				{
					"type": "struct",
					"name": "IndexAndValues",
					"size": 3,
					"members": [
						{
							"name": "index",
							"type": "felt",
							"offset": 0
						},
						{
							"name": "values",
							"type": "(felt, felt)",
							"offset": 1
						}
					]
				},
				{
					"type": "function",
					"name": "advance_counter",
					"inputs": [
						{
							"name": "index",
							"type": "felt"
						},
						{
							"name": "diffs_len",
							"type": "felt"
						},
						{
							"name": "diffs",
							"type": "felt*"
						}
					],
					"outputs": []
				},
				{
					"type": "constructor",
					"name": "constructor",
					"inputs": [
						{
							"name": "address",
							"type": "felt"
						},
						{
							"name": "value",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "xor_counters",
					"inputs": [
						{
							"name": "index_and_x",
							"type": "IndexAndValues"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "call_xor_counters",
					"inputs": [
						{
							"name": "address",
							"type": "felt"
						},
						{
							"name": "index_and_x",
							"type": "IndexAndValues"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "add_signature_to_counters",
					"inputs": [
						{
							"name": "index",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "set_value",
					"inputs": [
						{
							"name": "address",
							"type": "felt"
						},
						{
							"name": "value",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "get_value",
					"inputs": [
						{
							"name": "address",
							"type": "felt"
						}
					],
					"outputs": [
						{
							"name": "res",
							"type": "felt"
						}
					]
				},
				{
					"type": "function",
					"name": "entry_point",
					"inputs": [],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_builtins",
					"inputs": [],
					"outputs": [
						{
							"name": "result",
							"type": "felt"
						}
					]
				},
				{
					"type": "function",
					"name": "send_message",
					"inputs": [
						{
							"name": "to_address",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_call_contract",
					"inputs": [
						{
							"name": "contract_address",
							"type": "felt"
						},
						{
							"name": "function_selector",
							"type": "felt"
						},
						{
							"name": "calldata_len",
							"type": "felt"
						},
						{
							"name": "calldata",
							"type": "felt*"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_delegate_call",
					"inputs": [
						{
							"name": "contract_address",
							"type": "felt"
						},
						{
							"name": "function_selector",
							"type": "felt"
						},
						{
							"name": "calldata_len",
							"type": "felt"
						},
						{
							"name": "calldata",
							"type": "felt*"
						}
					],
					"outputs": []
				},
				{
					"type": "l1_handler",
					"name": "deposit",
					"inputs": [
						{
							"name": "from_address",
							"type": "felt"
						},
						{
							"name": "amount",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_get_caller_address",
					"inputs": [
						{
							"name": "expected_address",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_get_sequencer_address",
					"inputs": [
						{
							"name": "expected_address",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_get_contract_address",
					"inputs": [
						{
							"name": "expected_address",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_call_storage_consistency",
					"inputs": [
						{
							"name": "other_contract_address",
							"type": "felt"
						},
						{
							"name": "address",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_re_entrance",
					"inputs": [
						{
							"name": "other_contract_address",
							"type": "felt"
						},
						{
							"name": "depth",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "add_value",
					"inputs": [
						{
							"name": "value",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "recursive_add_value",
					"inputs": [
						{
							"name": "self_address",
							"type": "felt"
						},
						{
							"name": "value",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "increase_value",
					"inputs": [
						{
							"name": "address",
							"type": "felt"
						}
					],
					"outputs": []
				},
				{
					"type": "function",
					"name": "test_call_with_array",
					"inputs": [
						{
							"name": "self_address",
							"type": "felt"
						},
						{
							"name": "arr_len",
							"type": "felt"
						},
						{
							"name": "arr",
							"type": "felt*"
						}
					],
					"outputs": []
				}
			]`),
			want: Abi{
				Functions: map[string]*FunctionAbiItem{
					"advance_counter": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "advance_counter",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "index",
							}, {
								Type: "felt",
								Name: "diffs_len",
							}, {
								Type: "felt",
								Name: "diffs",
							},
						},
						Outputs: []Type{},
					},
					"xor_counters": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "xor_counters",
						},
						Inputs: []Type{
							{
								Type: "IndexAndValues",
								Name: "index_and_x",
							},
						},
						Outputs: []Type{},
					},
					"call_xor_counters": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "call_xor_counters",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "address",
							}, {
								Type: "IndexAndValues",
								Name: "index_and_x",
							},
						},
						Outputs: []Type{},
					},
					"add_signature_to_counters": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "add_signature_to_counters",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "index",
							},
						},
						Outputs: []Type{},
					},
					"set_value": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "set_value",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "address",
							}, {
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []Type{},
					},
					"get_value": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "get_value",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "address",
							},
						},
						Outputs: []Type{
							{
								Type: "felt",
								Name: "res",
							},
						},
					},
					"entry_point": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "entry_point",
						},
						Inputs:  []Type{},
						Outputs: []Type{},
					},
					"test_builtins": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_builtins",
						},
						Inputs: []Type{},
						Outputs: []Type{
							{
								Type: "felt",
								Name: "res",
							},
						},
					},
					"send_message": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "send_message",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "to_address",
							},
						},
						Outputs: []Type{},
					},
					"test_call_contract": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_call_contract",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "contract_address",
							}, {
								Type: "felt",
								Name: "function_selector",
							}, {
								Type: "felt",
								Name: "calldata_len",
							}, {
								Type: "felt",
								Name: "calldata",
							},
						},
						Outputs: []Type{},
					},
					"test_delegate_call": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_delegate_call",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "contract_address",
							}, {
								Type: "felt",
								Name: "function_selector",
							}, {
								Type: "felt",
								Name: "calldata_len",
							}, {
								Type: "felt",
								Name: "calldata",
							},
						},
						Outputs: []Type{},
					},
					"test_get_caller_address": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_get_caller_address",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "expected_address",
							},
						},
						Outputs: []Type{},
					},
					"test_get_sequencer_address": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_get_sequencer_address",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "expected_address",
							},
						},
						Outputs: []Type{},
					},
					"test_get_contract_address": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_get_contract_address",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "expected_address",
							},
						},
						Outputs: []Type{},
					},
					"test_call_storage_consistency": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_call_storage_consistency",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "other_contract_address",
							}, {
								Type: "felt",
								Name: "address",
							},
						},
						Outputs: []Type{},
					},
					"test_re_entrance": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_re_entrance",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "other_contract_address",
							}, {
								Type: "felt",
								Name: "depth",
							},
						},
						Outputs: []Type{},
					},
					"add_value": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "add_value",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []Type{},
					},
					"recursive_add_value": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "recursive_add_value",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "self_address",
							}, {
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []Type{},
					},
					"increase_value": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "increase_value",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "address",
							},
						},
						Outputs: []Type{},
					},
					"test_call_with_array": {
						Type: Type{
							Type: AbiFunctionType,
							Name: "test_call_with_array",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "self_address",
							}, {
								Type: "felt",
								Name: "arr_len",
							}, {
								Type: "felt",
								Name: "arr",
							},
						},
						Outputs: []Type{},
					},
				},
				L1Handlers: map[string]*FunctionAbiItem{
					"deposit": {
						Type: Type{
							Type: AbiL1HandlerType,
							Name: "deposit",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "from_address",
							}, {
								Type: "felt",
								Name: "amount",
							},
						},
						Outputs: []Type{},
					},
				},
				Constructor: map[string]*FunctionAbiItem{
					"constructor": {
						Type: Type{
							Type: AbiConstructorType,
							Name: "constructor",
						},
						Inputs: []Type{
							{
								Type: "felt",
								Name: "address",
							}, {
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []Type{},
					},
				},
				Events: map[string]*EventAbiItem{},
				Structs: map[string]*StructAbiItem{
					"IndexAndValues": {
						Type: Type{
							Type: AbiStructType,
							Name: "IndexAndValues",
						},
						Members: []Member{
							{
								Type: Type{
									Type: "felt",
									Name: "index",
								},
								Offset: 0,
							}, {
								Type: Type{
									Type: "(felt, felt)",
									Name: "values",
								},
								Offset: 0,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a Abi
			if err := a.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("Abi.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Len(t, a.Functions, len(tt.want.Functions), "invalid functions count") {
				t.Errorf("invalid functions count: want = %d got=%d", len(tt.want.Functions), len(a.Functions))
				return
			}
			if !assert.Len(t, a.L1Handlers, len(tt.want.L1Handlers), "invalid l1 handlers count") {
				t.Errorf("invalid l1_handlers count: want = %d got=%d", len(tt.want.L1Handlers), len(a.L1Handlers))
				return
			}
			if !assert.Len(t, a.Constructor, len(tt.want.Constructor), "invalid constructors count") {
				t.Errorf("invalid constructors count: want = %d got=%d", len(tt.want.Constructor), len(a.Constructor))
				return
			}
			if !assert.Len(t, a.Structs, len(tt.want.Structs), "invalid structs count") {
				t.Errorf("invalid structs count: want = %d got=%d", len(tt.want.Structs), len(a.Structs))
				return
			}
			if !assert.Len(t, a.Events, len(tt.want.Events), "invalid events count") {
				t.Errorf("invalid events count: want = %d got=%d", len(tt.want.Events), len(a.Events))
				return
			}
		})
	}
}
