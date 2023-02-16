package starknetgoapi

import "context"

// GetNonce -
func (api API) GetNonce(ctx context.Context, contract string, block BlockFilter, opts ...RequestOption) (*Response[string], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getNonce", []any{
		block, contract,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
