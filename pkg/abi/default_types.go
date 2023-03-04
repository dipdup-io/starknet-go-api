package abi

import "github.com/dipdup-io/starknet-go-api/pkg/encoding"

// default types
var (
	ExecuteFunction = FunctionItem{
		Type: Type{
			Type: FunctionType,
			Name: encoding.ExecuteEntrypoint,
		},

		Inputs: []Type{
			{
				Name: "call_array_len",
				Type: "felt",
			}, {
				Name: "call_array",
				Type: "CallArray*",
			}, {
				Name: "calldata_len",
				Type: "felt",
			}, {
				Name: "calldata",
				Type: "felt*",
			},
		},
		Outputs: []Type{
			{
				Name: "response_len",
				Type: "felt",
			}, {
				Name: "response",
				Type: "felt*",
			},
		},
	}

	CallArray = StructItem{
		Type: Type{
			Type: StructType,
			Name: "CallArray",
		},
		Size: 4,
		Members: []Member{
			{
				Type: Type{
					Type: "felt",
					Name: "to",
				},
				Offset: 0,
			}, {
				Type: Type{
					Type: "felt",
					Name: "selector",
				},
				Offset: 1,
			}, {
				Type: Type{
					Type: "felt",
					Name: "data_offset",
				},
				Offset: 2,
			}, {
				Type: Type{
					Type: "felt",
					Name: "data_len",
				},
				Offset: 3,
			},
		},
	}
)
