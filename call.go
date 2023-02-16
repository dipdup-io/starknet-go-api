package starknetgoapi

import "context"

// Call - Calls a function in a contract and returns the return value. Using this call will not create a transaction; hence, will not change the state
func (api API) Call(ctx context.Context, params CallRequest, block BlockFilter, opts ...RequestOption) (*Response[[]string], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_call", []any{
		params, block,
	}, opts...)

	var response Response[[]string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
