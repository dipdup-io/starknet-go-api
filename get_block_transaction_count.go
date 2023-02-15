package starknetgoapi

import "context"

// GetTransactionCountByBlockNumber -
func (api API) GetTransactionCountByBlockNumber(ctx context.Context, blockNumber uint64, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockTransactionCount", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
	}, opts...)

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetTransactionCountByBlockHash -
func (api API) GetTransactionCountByBlockHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockTransactionCount", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
	}, opts...)

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetTransactionCountByLatestBlock -
func (api API) GetTransactionCountByLatestBlock(ctx context.Context, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockTransactionCount", []any{
		latest,
	}, opts...)

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetTransactionCountByPendingBlock -
func (api API) GetTransactionCountByPendingBlock(ctx context.Context, opts ...RequestOption) (*Response[uint64], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockTransactionCount", []any{
		pending,
	}, opts...)

	var response Response[uint64]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
