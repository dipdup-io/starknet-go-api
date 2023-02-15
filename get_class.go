package starknetgoapi

import "context"

// Class -
type Class struct {
	Program           string            `json:"program"`
	EntryPointsByType EntrypointsByType `json:"entry_points_by_type"`
	Abi               []Abi             `json:"abi"`
}

// Handler -
type Handler struct {
	Offset   string `json:"offset"`
	Selector string `json:"selector"`
}

// Type -
type Type struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Size   uint64 `json:"size,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}

// Abi -
type Abi struct {
	Type
	Members []Type   `json:"members,omitempty"`
	Inputs  []Type   `json:"inputs,omitempty"`
	Outputs []Type   `json:"outputs,omitempty"`
	Data    []Type   `json:"data,omitempty"`
	Keys    []string `json:"keys,omitempty"`
}

// EntrypointsByType -
type EntrypointsByType struct {
	CONSTRUCTOR []Handler `json:"CONSTRUCTOR"`
	EXTERNAL    []Handler `json:"EXTERNAL"`
	L1HANDLER   []Handler `json:"L1_HANDLER"`
}

// GetClass -
func (api API) GetClass(ctx context.Context, block BlockFilter, classHash string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClass", []any{
		block, classHash,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
