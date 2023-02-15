package starknetgoapi

import "context"

// GetBlockTransactionCount -
func (api API) GetBlockTransactionCount(ctx context.Context, filters BlockFilter, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockTransactionCount", []any{filters}, opts...)

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
