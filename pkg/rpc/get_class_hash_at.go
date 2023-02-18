package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetClassHashAt -
func (api API) GetClassHashAt(ctx context.Context, block data.BlockID, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		block, contractAddress,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
