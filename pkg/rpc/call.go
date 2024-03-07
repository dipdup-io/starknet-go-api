package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// Call - Calls a function in a contract and returns the return value. Using this call will not create a transaction; hence, will not change the state
func (api API) Call(ctx context.Context, params CallRequest, block data.BlockID, opts ...RequestOption) (*Response[[]string], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest("starknet_call", []any{
		params, block,
	}, opts...)

	var response Response[[]string]
	err := post(ctx, api, *request, &response)
	return &response, err
}
