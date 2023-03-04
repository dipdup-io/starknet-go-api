package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetStateUpdate -
func (api API) GetStateUpdate(ctx context.Context, block data.BlockID, opts ...RequestOption) (*Response[data.StateUpdate], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{block}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[data.StateUpdate]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
