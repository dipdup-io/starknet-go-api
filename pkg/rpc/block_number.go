package api

import "context"

// BlockNumber - Get the most recent accepted block number
func (api API) BlockNumber(ctx context.Context, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest("starknet_blockNumber", []any{}, opts...)

	var response Response[uint64]
	err := post(ctx, api, *request, &response)
	return &response, err
}
