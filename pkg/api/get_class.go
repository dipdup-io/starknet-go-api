package api

import "context"

// GetClass - Get the contract class definition in the given block associated with the given hash
func (api API) GetClass(ctx context.Context, block BlockFilter, classHash string, opts ...RequestOption) (*Response[Class], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getClass", []any{
		block, classHash,
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
