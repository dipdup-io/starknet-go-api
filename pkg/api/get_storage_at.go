package api

import "context"

// GetStorageAt -
func (api API) GetStorageAt(ctx context.Context, contract, key string, block BlockFilter, opts ...RequestOption) (*Response[string], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		contract,
		key,
		block,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
