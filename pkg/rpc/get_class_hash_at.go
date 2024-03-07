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

	request := api.prepareRequest("starknet_getClassHashAt", []any{
		block, contractAddress,
	}, opts...)

	var response Response[string]
	err := post(ctx, api, *request, &response)
	return &response, err
}
