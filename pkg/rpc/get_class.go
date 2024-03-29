package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetClass - Get the contract class definition in the given block associated with the given hash
func (api API) GetClass(ctx context.Context, block data.BlockID, classHash string, opts ...RequestOption) (*Response[data.Class], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest("starknet_getClass", []any{
		block, classHash,
	}, opts...)

	var response Response[data.Class]
	err := post(ctx, api, *request, &response)
	return &response, err
}
