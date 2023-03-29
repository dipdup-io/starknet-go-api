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
