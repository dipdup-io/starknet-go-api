package api

import "context"

// GetClassHashAt -
func (api API) GetClassHashAt(ctx context.Context, block BlockFilter, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		block, contractAddress,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
