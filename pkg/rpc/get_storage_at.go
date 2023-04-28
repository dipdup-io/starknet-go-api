package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetStorageAt -
func (api API) GetStorageAt(ctx context.Context, contract, key string, block data.BlockID, opts ...RequestOption) (*Response[string], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		contract,
		key,
		block,
	}, opts...)

	var response Response[string]
	err := post(ctx, api, *request, &response)
	return &response, err
}
