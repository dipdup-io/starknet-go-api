package abi

import (
	"reflect"
	"testing"

	"github.com/goccy/go-json"
)

func TestDecodeFunctionCallData(t *testing.T) {
	type args struct {
		calldata []string
		abi      string
		endpoint string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				calldata: []string{
					"0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
					"0x0",
				},
				endpoint: "test_re_entrance",
				abi: `[{
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
						}]`,
			},
			want: map[string]any{
				"other_contract_address": "0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
				"depth":                  "0x0",
			},
		}, {
			name: "test 2",
			args: args{
				calldata: []string{
					"0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
					"0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
					"0x4",
					"0x5bd24b507fcc2fd77dc7847babb8df01363d58e9b0bbcd2d06d982e1f3e0c86",
					"0x2",
					"0x26b5943d4a0c420607cee8030a8cdd859bf2814a06633d165820960a42c6aed",
					"0x1518eec76afd5397cefd14eda48d01ad59981f9ce9e70c233ca67acd8754008",
				},
				endpoint: "test_call_contract",
				abi: `[{
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
					}]`,
			},
			want: map[string]any{
				"contract_address":  "0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
				"function_selector": "0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
				"calldata_len":      "0x4",
				"calldata": []any{
					"0x5bd24b507fcc2fd77dc7847babb8df01363d58e9b0bbcd2d06d982e1f3e0c86",
					"0x2",
					"0x26b5943d4a0c420607cee8030a8cdd859bf2814a06633d165820960a42c6aed",
					"0x1518eec76afd5397cefd14eda48d01ad59981f9ce9e70c233ca67acd8754008",
				},
			},
		}, {
			name: "test 3",
			args: args{
				calldata: []string{
					"0xf899bfebffcf95f54cf61ac07a85587988678ed4",
					"0xde0b6b3a7640000",
					"0x0",
				},
				endpoint: "initiate_withdraw",
				abi: `[{
						"type": "function",
						"name": "initiate_withdraw",
						"inputs": [
						  {
							"name": "l1_recipient",
							"type": "felt"
						  },
						  {
							"name": "amount",
							"type": "Uint256"
						  }
						],
						"outputs": []
					},{
						"name": "Uint256",
						"size": 2,
						"members": [
						  {
							"name": "low",
							"type": "felt",
							"offset": 0
						  },
						  {
							"name": "high",
							"type": "felt",
							"offset": 1
						  }
						],
						"type": "struct"
					}]`,
			},
			want: map[string]any{
				"l1_recipient": "0xf899bfebffcf95f54cf61ac07a85587988678ed4",
				"amount": map[string]any{
					"low":  "0xde0b6b3a7640000",
					"high": "0x0",
				},
			},
		}, {
			name: "test 4",
			args: args{
				calldata: []string{
					"0x1",
					"0x41a78e741e5af2fec34b695679bc6891742439f7afb8484ecd7766661ad02bf",
					"0x1987cbd17808b9a23693d4de7e246a443cfe37e6e7fbaeabd7d7e6532b07c3d",
					"0xa",
					"0x48e09e58a43b0794821ac868ed3c2d940a2ff13a2a7edc57dea999cf881e439",
					"0x52554c45532056322e30",
					"0x0",
					"0x6",
					"0x2",
					"0x68747470733a2f2f6170692e72756c65732e6172742f6d657461646174612f",
					"0x7b69647d2e6a736f6e",
					"0x7ca95733f27cbd125214ada1c131a100daca25354cb9bfcac1ef5ea479646b5",
					"0x4b5cab01f7e6e59df39c14a3f3436c721554f719e42618fd4e69e4a605d39ef",
					"0x63a4b3b0122cdaa6ba244739add94aed1d31e3330458cda833a8d119f28cbe8",
				},
				endpoint: "__execute__",
				abi: `[{
					"name": "rules_account::account::interface::Call",
					"type": "struct",
					"members": [
					  {
						"name": "to",
						"type": "core::starknet::contract_address::ContractAddress"
					  },
					  {
						"name": "selector",
						"type": "core::felt252"
					  },
					  {
						"name": "calldata",
						"type": "core::array::Array::<core::felt252>"
					  }
					]
				},{
					"name": "__execute__",
					"type": "function",
					"inputs": [
					  {
						"name": "calls",
						"type": "core::array::Array::<rules_account::account::interface::Call>"
					  }
					],
					"outputs": [
					  {
						"type": "core::array::Array::<core::array::Span::<core::felt252>>"
					  }
					],
					"state_mutability": "external"
				}]`,
			},
			want: map[string]any{
				"calls": []any{
					map[string]any{
						"selector": "0x1987cbd17808b9a23693d4de7e246a443cfe37e6e7fbaeabd7d7e6532b07c3d",
						"to":       "0x41a78e741e5af2fec34b695679bc6891742439f7afb8484ecd7766661ad02bf",
						"calldata": []any{
							"0x48e09e58a43b0794821ac868ed3c2d940a2ff13a2a7edc57dea999cf881e439",
							"0x52554c45532056322e30",
							"0x0",
							"0x6",
							"0x2",
							"0x68747470733a2f2f6170692e72756c65732e6172742f6d657461646174612f",
							"0x7b69647d2e6a736f6e",
							"0x7ca95733f27cbd125214ada1c131a100daca25354cb9bfcac1ef5ea479646b5",
							"0x4b5cab01f7e6e59df39c14a3f3436c721554f719e42618fd4e69e4a605d39ef",
							"0x63a4b3b0122cdaa6ba244739add94aed1d31e3330458cda833a8d119f28cbe8",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a Abi
			if err := json.Unmarshal([]byte(tt.args.abi), &a); err != nil {
				t.Errorf("DecodeFunctionCallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := DecodeFunctionCallData(tt.args.calldata, *a.Functions[tt.args.endpoint], a.Structs)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFunctionCallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeFunctionCallData() = %v, want %v", got, tt.want)
			}
		})
	}
}
