package starknetgoapi

import "context"

// GetNonce -
func (api API) GetNonce(ctx context.Context, contract string, block BlockFilter, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getNonce", []any{
		block, contract,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
