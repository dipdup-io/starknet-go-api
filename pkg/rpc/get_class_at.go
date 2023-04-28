package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetClassAt -
func (api API) GetClassAt(ctx context.Context, block data.BlockID, contractAddress string, opts ...RequestOption) (*Response[data.Class], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		block, contractAddress,
	}, opts...)

	var response Response[data.Class]
	err := post(ctx, api, *request, &response)
	return &response, err
}
