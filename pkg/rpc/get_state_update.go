package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetStateUpdate -
func (api API) GetStateUpdate(ctx context.Context, block data.BlockID, opts ...RequestOption) (*Response[data.StateUpdateRpc], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{block}, opts...)

	var response Response[data.StateUpdateRpc]
	err := post(ctx, api, *request, &response)
	return &response, err
}
