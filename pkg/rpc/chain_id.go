package api

import "context"

// ChainID - Return the currently configured StarkNet chain id
func (api API) ChainID(ctx context.Context, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest("starknet_chainId", []any{}, opts...)

	var response Response[string]
	err := post(ctx, api, *request, &response)
	return &response, err
}
