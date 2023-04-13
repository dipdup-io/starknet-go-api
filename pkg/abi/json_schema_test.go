package abi

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/tfkhsr/jsonschema"
)

func TestAbi_JsonSchema(t *testing.T) {
	tests := []struct {
		name string
		abi  []byte
	}{
		{
			name: "test 1",
			abi: []byte(`[
				{
				  "name": "Uint256",
				  "size": 2,
				  "type": "struct",
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
				  ]
				},
				{
				  "data": [
					{
					  "name": "from_",
					  "type": "felt"
					},
					{
					  "name": "to",
					  "type": "felt"
					},
					{
					  "name": "value",
					  "type": "Uint256"
					}
				  ],
				  "keys": [],
				  "name": "Transfer",
				  "type": "event"
				},
				{
				  "data": [
					{
					  "name": "owner",
					  "type": "felt"
					},
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "value",
					  "type": "Uint256"
					}
				  ],
				  "keys": [],
				  "name": "Approval",
				  "type": "event"
				},
				{
				  "name": "constructor",
				  "type": "constructor",
				  "inputs": [
					{
					  "name": "recipient",
					  "type": "felt"
					}
				  ],
				  "outputs": []
				},
				{
				  "name": "name",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "name",
					  "type": "felt"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "symbol",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "symbol",
					  "type": "felt"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "totalSupply",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "totalSupply",
					  "type": "Uint256"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "decimals",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "decimals",
					  "type": "felt"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "balanceOf",
				  "type": "function",
				  "inputs": [
					{
					  "name": "account",
					  "type": "felt"
					}
				  ],
				  "outputs": [
					{
					  "name": "balance",
					  "type": "Uint256"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "allowance",
				  "type": "function",
				  "inputs": [
					{
					  "name": "owner",
					  "type": "felt"
					},
					{
					  "name": "spender",
					  "type": "felt"
					}
				  ],
				  "outputs": [
					{
					  "name": "remaining",
					  "type": "Uint256"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "transfer",
				  "type": "function",
				  "inputs": [
					{
					  "name": "recipient",
					  "type": "felt"
					},
					{
					  "name": "amount",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "transferFrom",
				  "type": "function",
				  "inputs": [
					{
					  "name": "sender",
					  "type": "felt"
					},
					{
					  "name": "recipient",
					  "type": "felt"
					},
					{
					  "name": "amount",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "approve",
				  "type": "function",
				  "inputs": [
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "amount",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "increaseAllowance",
				  "type": "function",
				  "inputs": [
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "added_value",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "decreaseAllowance",
				  "type": "function",
				  "inputs": [
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "subtracted_value",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				}
			  ]`),
		},
		{
			name: "test 2",
			abi:  []byte(`[{"type": "event", "name": "Hello", "keys": [{"name": "from", "type": "core::starknet::contract_address::ContractAddress"}, {"name": "value", "type": "core::integer::u128"}]}]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var abi Abi
			if err := json.Unmarshal(tt.abi, &abi); err != nil {
				t.Errorf("can't unmarshal abi: %s", string(tt.abi))
				return
			}
			got := abi.JsonSchema()
			b, err := json.MarshalIndent(got, "", " ")
			if err != nil {
				t.Errorf("can't marshal jsonschema: %s", string(tt.abi))
				return
			}

			if _, err := jsonschema.Parse(b); err != nil {
				t.Errorf("can't parse resulting json schema: %s", string(tt.abi))
			}

		})
	}
}
