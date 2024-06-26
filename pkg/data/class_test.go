package data

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/dipdup-io/starknet-go-api/pkg/abi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAbi_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    abi.Abi
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
			want: abi.Abi{
				Functions: map[string]*abi.FunctionItem{
					"advance_counter": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "advance_counter",
						},
						Inputs: []abi.Type{
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
						Outputs: []abi.Type{},
					},
					"xor_counters": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "xor_counters",
						},
						Inputs: []abi.Type{
							{
								Type: "IndexAndValues",
								Name: "index_and_x",
							},
						},
						Outputs: []abi.Type{},
					},
					"call_xor_counters": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "call_xor_counters",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "address",
							}, {
								Type: "IndexAndValues",
								Name: "index_and_x",
							},
						},
						Outputs: []abi.Type{},
					},
					"add_signature_to_counters": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "add_signature_to_counters",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "index",
							},
						},
						Outputs: []abi.Type{},
					},
					"set_value": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "set_value",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "address",
							}, {
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []abi.Type{},
					},
					"get_value": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "get_value",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "address",
							},
						},
						Outputs: []abi.Type{
							{
								Type: "felt",
								Name: "res",
							},
						},
					},
					"entry_point": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "entry_point",
						},
						Inputs:  []abi.Type{},
						Outputs: []abi.Type{},
					},
					"test_builtins": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_builtins",
						},
						Inputs: []abi.Type{},
						Outputs: []abi.Type{
							{
								Type: "felt",
								Name: "res",
							},
						},
					},
					"send_message": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "send_message",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "to_address",
							},
						},
						Outputs: []abi.Type{},
					},
					"test_call_contract": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_call_contract",
						},
						Inputs: []abi.Type{
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
						Outputs: []abi.Type{},
					},
					"test_delegate_call": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_delegate_call",
						},
						Inputs: []abi.Type{
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
						Outputs: []abi.Type{},
					},
					"test_get_caller_address": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_get_caller_address",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "expected_address",
							},
						},
						Outputs: []abi.Type{},
					},
					"test_get_sequencer_address": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_get_sequencer_address",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "expected_address",
							},
						},
						Outputs: []abi.Type{},
					},
					"test_get_contract_address": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_get_contract_address",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "expected_address",
							},
						},
						Outputs: []abi.Type{},
					},
					"test_call_storage_consistency": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_call_storage_consistency",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "other_contract_address",
							}, {
								Type: "felt",
								Name: "address",
							},
						},
						Outputs: []abi.Type{},
					},
					"test_re_entrance": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_re_entrance",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "other_contract_address",
							}, {
								Type: "felt",
								Name: "depth",
							},
						},
						Outputs: []abi.Type{},
					},
					"add_value": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "add_value",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []abi.Type{},
					},
					"recursive_add_value": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "recursive_add_value",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "self_address",
							}, {
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []abi.Type{},
					},
					"increase_value": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "increase_value",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "address",
							},
						},
						Outputs: []abi.Type{},
					},
					"test_call_with_array": {
						Type: abi.Type{
							Type: abi.FunctionType,
							Name: "test_call_with_array",
						},
						Inputs: []abi.Type{
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
						Outputs: []abi.Type{},
					},
				},
				L1Handlers: map[string]*abi.FunctionItem{
					"deposit": {
						Type: abi.Type{
							Type: abi.L1HandlerType,
							Name: "deposit",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "from_address",
							}, {
								Type: "felt",
								Name: "amount",
							},
						},
						Outputs: []abi.Type{},
					},
				},
				Constructor: map[string]*abi.FunctionItem{
					"constructor": {
						Type: abi.Type{
							Type: abi.ConstructorType,
							Name: "constructor",
						},
						Inputs: []abi.Type{
							{
								Type: "felt",
								Name: "address",
							}, {
								Type: "felt",
								Name: "value",
							},
						},
						Outputs: []abi.Type{},
					},
				},
				Events: map[string]*abi.EventItem{},
				Structs: map[string]*abi.StructItem{
					"IndexAndValues": {
						Type: abi.Type{
							Type: abi.StructType,
							Name: "IndexAndValues",
						},
						Members: []abi.Member{
							{
								Type: abi.Type{
									Type: "felt",
									Name: "index",
								},
								Offset: 0,
							}, {
								Type: abi.Type{
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
			var a abi.Abi
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

func TestClass_GetAbi(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{
			name: "test 1",
			data: []byte(`{"contract_class_version": "0.1.0", "sierra_program": ["0x1", "0x5", "0x0", "0x2", "0x6", "0x0", "0x97", "0x69", "0x18", "0x52616e6765436865636b", "0x800000000000000100000000000000000000000000000000", "0x436f6e7374", "0x800000000000000000000000000000000000000000000002", "0x1", "0xc", "0x2", "0x4661696c656420746f20646573657269616c697a6520706172616d202331", "0x4f7574206f6620676173", "0x4172726179", "0x800000000000000300000000000000000000000000000001", "0x536e617073686f74", "0x800000000000000700000000000000000000000000000001", "0x3", "0x537472756374", "0x800000000000000700000000000000000000000000000002", "0x0", "0x1baeba72e79e9db2587cf44fedb2f3700b2075a5e8e39a562584862c4b71f62", "0x4", "0x2ee1e2b1b89f8c495f200e4956278a4d47395fe262f27b52e5865c9524c08c3", "0x5", "0x8", "0x753332", "0x800000000000000700000000000000000000000000000000", "0x53746f7261676541646472657373", "0x53746f726167654261736541646472657373", "0x456e7465722061206e616d65", "0x66656c74323532", "0x4e6f6e5a65726f", "0x4275696c74696e436f737473", "0x53797374656d", "0x800000000000000f00000000000000000000000000000001", "0x16a4c8d7c05909052238a862d8cc3e7975bf05a07b3a69c6b28951083a6d672", "0x800000000000000300000000000000000000000000000003", "0x10", "0x456e756d", "0x9931c641b913035ae674b400b61a51476d506bbe8bba2ff8a6272790aba9e6", "0x6", "0x11", "0x496e70757420746f6f206c6f6e6720666f7220617267756d656e7473", "0x800000000000000700000000000000000000000000000003", "0x11c6d8087e00642489f92d2821ad6ebd6532ad1a3b6d12833da6d6810391511", "0x14", "0x426f78", "0x4761734275696c74696e", "0x36", "0x7265766f6b655f61705f747261636b696e67", "0x77697468647261775f676173", "0x6272616e63685f616c69676e", "0x7374727563745f6465636f6e737472756374", "0x656e61626c655f61705f747261636b696e67", "0x73746f72655f74656d70", "0x61727261795f736e617073686f745f706f705f66726f6e74", "0x756e626f78", "0x72656e616d65", "0x656e756d5f696e6974", "0x15", "0x6a756d70", "0x7374727563745f636f6e737472756374", "0x656e756d5f6d61746368", "0x64697361626c655f61705f747261636b696e67", "0x64726f70", "0x16", "0x61727261795f6e6577", "0x636f6e73745f61735f696d6d656469617465", "0x13", "0x61727261795f617070656e64", "0x12", "0x17", "0xf", "0x6765745f6275696c74696e5f636f737473", "0xe", "0x77697468647261775f6761735f616c6c", "0x647570", "0x66656c743235325f69735f7a65726f", "0xb", "0xd", "0x73746f726167655f626173655f616464726573735f636f6e7374", "0x361458367e696363fbcc70777d07ebbd2394e89fd0adcaf147faccd1d294d60", "0x73746f726167655f616464726573735f66726f6d5f62617365", "0x7", "0x9", "0x73746f726167655f77726974655f73797363616c6c", "0x736e617073686f745f74616b65", "0x73746f726167655f726561645f73797363616c6c", "0xde", "0xffffffffffffffff", "0x7c", "0xa", "0x6c", "0x27", "0x19", "0x1a", "0x1b", "0x1c", "0x1d", "0x1e", "0x1f", "0x20", "0x5e", "0x21", "0x22", "0x23", "0x3c", "0x24", "0x25", "0x26", "0x28", "0x29", "0x2a", "0x55", "0x2b", "0x2c", "0x2d", "0x2e", "0x2f", "0x51", "0x30", "0x31", "0x32", "0x33", "0x34", "0x35", "0x37", "0x38", "0x39", "0x3a", "0x3b", "0x3d", "0x3e", "0x3f", "0x40", "0x41", "0x42", "0x43", "0x44", "0x45", "0x46", "0x47", "0x48", "0x49", "0x4a", "0x4b", "0x4c", "0x4d", "0xd0", "0x9f", "0xc3", "0xba", "0x8a", "0x80e", "0x100f13051211100f0e050d0c06050b0a090706050403080706050403020100", "0x1f181e06050d1d181c1b0706050403181a1819181711071605040315051411", "0x1411200f28070605040327052605251122240e0523051411220f2111200f02", "0x1105053411331505053211311130112f2e022d06050d2c2b0506052a112924", "0x5053b113a3905053413050534110739050738060505370605053606050535", "0x541060505400605053e3f05053e1305053e113d3905053c0507390507382b", "0x50534450505340507440507382705053b2305053b06050543060505344205", "0x4d05053e0e0505344c050541114b0605054a11494805053411474405053446", "0x5053b0e05053e0e050554115352050534160505345105054111504f07054e", "0x115611551505053e050505412b05053e070505411107440507382605053b15", "0x57050e050e1111570511071151260758151307570705110705111157051111", "0x1107114c054859520757071605261113055705130515111157051113111605", "0x52055911480557054d0552114d055705060516110605570559055111115705", "0x27055705114d111157051107111142051106112305570548054c1146055705", "0x55707230546112305570544054c11460557054c0559114405570527054811", "0x11231111570511071145055b3f39075707460526111157051107112b055a42", "0x55705112b11115705420542111157053f05441111570539052711115705", "0x5e0557051100115d0557055c000745115c0557055c053f115c055705113911", "0x570515055e111305570513051511600557055f055d115f0557055d5e075c11", "0x1111570511071160071513130560055705600560110705570507055f111505", "0x130e6311610557056105621161055705116111115705450527111157051123", "0x42075705420565111157051113111157051107116665076463620757076115", "0x57054205421111570511071169056811570767056611620557056205151167", "0x557056b6a0745116b0557056b053f116b0557051167116a055705112b1111", "0x116f051106116e0557055b0569116d05570507055f116c05570563055e115b", "0x116c117105570570055b1170055705116b1111570569056a11115705110711", "0x7570742717207631570117105570571056e117205570572056d1172055705", "0x57905711179055705112b111157051123111157051107117877760e757473", "0x7c0576117c0557052e0574112e0557057b0573111157057a0572117b7a0757", "0x560117405570574055f117305570573055e1162055705620515117d055705", "0x570577055f116c05570576055e111157051107117d74736213057d0557057d", "0x557056e7e075c117e0557051100111157051123116e055705780569116d05", "0x57056d055f116c0557056c055e116205570562051511800557057f055d117f", "0x2b1111570542054211115705110711806d6c62130580055705800560116d05", "0x118205570581680745118105570581053f11810557051177116805570511", "0x1165055705650515118505570584055d11840557058283075c118305570511", "0x71185076665130585055705850560110705570507055f116605570566055e", "0x86055705112b11115705460527111157052b05781111570511231111570511", "0x890557051100118805570587860745118705570587053f1187055705117911", "0x570515055e1113055705130515118b0557058a055d118a0557058889075c11", "0x111157051107118b07151313058b0557058b0560110705570507055f111505", "0x745118d0557058d053f118d0557051177118c055705112b111157050e057a", "0x119105570590055d11900557058e8f075c118f0557051100118e0557058d8c", "0x91055705910560110705570507055f115105570551055e1126055705260515", "0x57051107115126079215130757070511070511115705111111910751261305", "0x7114c05935952075707160526111305570513051511160557050e050e1111", "0x5570511391106055705112b11115705590544111157055205271111570511", "0x54846075c1146055705110011480557054d060745114d0557054d053f114d", "0x7055f111505570515055e1113055705130515112705570523055d11230557", "0x1157054c052711115705110711270715131305270557052705601107055705", "0x113f3907942b420757074415130e6311440557054405621144055705116111", "0x5c056d115c055705116c110005570545055b1145055705116b111157051107", "0x5d0e5707005c072b137b1142055705420515110005570500056e115c055705", "0x745115f0557055f053f1163055705112b111157051107116261600e955f5e", "0x11690557056705731111570566057211676607570565057111650557055f63", "0x5d0557055d055e1142055705420515116b0557056a0576116a055705690574", "0x1100111157051107116b5e5d4213056b0557056b0560115e0557055e055f11", "0x5e1142055705420515116d0557056c055d116c055705625b075c115b055705", "0x1107116d61604213056d0557056d0560116105570561055f11600557056005", "0x5706e0745117005570570053f11700557051177116e055705112b11115705", "0x390515117405570573055d11730557057172075c1172055705110011710557", "0x39130574055705740560110705570507055f113f0557053f055e1139055705", "0x117705570511771176055705112b111157050e057a1111570511071174073f", "0x557057879075c11790557051100117805570577760745117705570577053f", "0x570507055f115105570551055e1126055705260515117b0557057a055d117a", "0xe07051144464511131546451113077b07512613057b0557057b0560110705", "0x960e0705114446451113154645111311"], "entry_points_by_type": {"CONSTRUCTOR": [], "EXTERNAL": [{"selector": "0xf61980aeb34c9c7f823d576c10d00648fdab6c03a59b539ed0824be31da466", "function_idx": 0}, {"selector": "0x1a07d8466becfa1c870148e061dc86efcdc4dc277243e9bbec3ac5bce9df28b", "function_idx": 1}], "L1_HANDLER": []}, "abi": "\"[{\\\"type\\\":\\\"impl\\\",\\\"name\\\":\\\"HelloStarknetImpl\\\",\\\"interface_name\\\":\\\"hello::hello::IHelloStarknet\\\"},{\\\"type\\\":\\\"interface\\\",\\\"name\\\":\\\"hello::hello::IHelloStarknet\\\",\\\"items\\\":[{\\\"type\\\":\\\"function\\\",\\\"name\\\":\\\"set_name\\\",\\\"inputs\\\":[{\\\"name\\\":\\\"name1\\\",\\\"type\\\":\\\"core::felt252\\\"}],\\\"outputs\\\":[],\\\"state_mutability\\\":\\\"external\\\"},{\\\"type\\\":\\\"function\\\",\\\"name\\\":\\\"get_name0000\\\",\\\"inputs\\\":[],\\\"outputs\\\":[{\\\"type\\\":\\\"core::felt252\\\"}],\\\"state_mutability\\\":\\\"view\\\"}]},{\\\"type\\\":\\\"event\\\",\\\"name\\\":\\\"hello::hello::HelloStarknet::Event\\\",\\\"kind\\\":\\\"enum\\\",\\\"variants\\\":[]}]\""}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Class
			err := json.Unmarshal(tt.data, &c)
			require.NoError(t, err)

			got, err := c.GetAbi()
			require.NoError(t, err)

			log.Println(got)
		})
	}
}
