package api

import "context"

// BlockNumber - Get the most recent accepted block number
func (api API) BlockNumber(ctx context.Context, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest(ctx, "starknet_blockNumber", []any{}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
