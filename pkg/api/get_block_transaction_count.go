package api

import "context"

// GetBlockTransactionCount -
func (api API) GetBlockTransactionCount(ctx context.Context, block BlockFilter, opts ...RequestOption) (*Response[uint64], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getBlockTransactionCount", []any{block}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
