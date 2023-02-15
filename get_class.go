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
	Members []Type `json:"members,omitempty"`
	Inputs  []Type `json:"inputs,omitempty"`
	Outputs []Type `json:"outputs,omitempty"`
}

// EntrypointsByType -
type EntrypointsByType struct {
	CONSTRUCTOR []Handler `json:"CONSTRUCTOR"`
	EXTERNAL    []Handler `json:"EXTERNAL"`
	L1HANDLER   []Handler `json:"L1_HANDLER"`
}

// GetClassByBlockNumber -
func (api API) GetClassByBlockNumber(ctx context.Context, blockNumber uint64, classHash string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClass", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
		classHash,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassByBlockHash -
func (api API) GetClassByBlockHash(ctx context.Context, hash, classHash string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClass", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
		classHash,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassByLatestBlock -
func (api API) GetClassByLatestBlock(ctx context.Context, classHash string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClass", []any{
		latest,
		classHash,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassByPendingBlock -
func (api API) GetClassByPendingBlock(ctx context.Context, classHash string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClass", []any{
		pending,
		classHash,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
