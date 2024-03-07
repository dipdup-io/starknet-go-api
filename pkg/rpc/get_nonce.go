package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetNonce -
func (api API) GetNonce(ctx context.Context, contract string, block data.BlockID, opts ...RequestOption) (*Response[string], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest("starknet_getNonce", []any{
		block, contract,
	}, opts...)

	var response Response[string]
	err := post(ctx, api, *request, &response)
	return &response, err
}
