package abi

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestEventItem_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		want EventItem
		data []byte
	}{
		{
			name: "new Transfer",
			data: []byte(`{
				"kind": "struct",
				"name": "openzeppelin::token::erc20_v070::erc20::ERC20::Transfer",
				"type": "event",
				"members": [
				{
					"kind": "data",
					"name": "from",
					"type": "core::starknet::contract_address::ContractAddress"
				},
				{
					"kind": "data",
					"name": "to",
					"type": "core::starknet::contract_address::ContractAddress"
				},
				{
					"kind": "data",
					"name": "value",
					"type": "core::integer::u256"
				}
				]
			}`),
			want: EventItem{
				Type: Type{
					Name: "openzeppelin::token::erc20_v070::erc20::ERC20::Transfer",
					Type: "event",
					Kind: "struct",
				},
				Data: []Type{
					{
						Kind: "data",
						Name: "from",
						Type: "core::starknet::contract_address::ContractAddress",
					}, {
						Kind: "data",
						Name: "to",
						Type: "core::starknet::contract_address::ContractAddress",
					}, {
						Kind: "data",
						Name: "value",
						Type: "core::integer::u256",
					},
				},
				Keys: []Type{},
			},
		}, {
			name: "old Transfer",
			data: []byte(`{
				"name": "transfer",
				"type": "event",
				"data": [
				{
					"name": "from",
					"type": "core::starknet::contract_address::ContractAddress"
				},
				{
					"name": "to",
					"type": "core::starknet::contract_address::ContractAddress"
				},
				{
					"name": "value",
					"type": "core::integer::u256"
				}
				],
				"keys":[]
			}`),
			want: EventItem{
				Type: Type{
					Name: "transfer",
					Type: "event",
				},
				Data: []Type{
					{
						Name: "from",
						Type: "core::starknet::contract_address::ContractAddress",
					}, {
						Name: "to",
						Type: "core::starknet::contract_address::ContractAddress",
					}, {
						Name: "value",
						Type: "core::integer::u256",
					},
				},
				Keys: []Type{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var item EventItem
			err := json.Unmarshal(tt.data, &item)
			require.NoError(t, err)
			require.Equal(t, tt.want, item)
		})
	}
}
