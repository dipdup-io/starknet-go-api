package starknetgoapi

import "context"

// GetClassHashAt -
func (api API) GetClassHashAt(ctx context.Context, block BlockFilter, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		block, contractAddress,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
