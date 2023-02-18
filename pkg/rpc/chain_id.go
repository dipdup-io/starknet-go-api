package api

import "context"

// ChainID - Return the currently configured StarkNet chain id
func (api API) ChainID(ctx context.Context, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_chainId", []any{}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
