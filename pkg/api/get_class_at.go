package api

import "context"

// GetClassAt -
func (api API) GetClassAt(ctx context.Context, block BlockFilter, contractAddress string, opts ...RequestOption) (*Response[Class], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		block, contractAddress,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
