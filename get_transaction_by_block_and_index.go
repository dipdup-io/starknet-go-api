package starknetgoapi

import "context"

// GetTransactionByBlockNumberAndIndex -
func (api API) GetTransactionByBlockNumberAndIndex(ctx context.Context, blockNumber, index uint64, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByHash", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
		index,
	}, opts...)

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetTransactionByBlockHashAndIndex -
func (api API) GetTransactionByBlockHashAndIndex(ctx context.Context, hash string, index uint64, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByHash", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
		index,
	}, opts...)

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetTransactionByLatestBlockAndIndex -
func (api API) GetTransactionByLatestBlockAndIndex(ctx context.Context, index uint64, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByHash", []any{
		latest,
		index,
	}, opts...)

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetTransactionByPendingBlockAndIndex -
func (api API) GetTransactionByPendingBlockAndIndex(ctx context.Context, index uint64, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByHash", []any{
		pending,
		index,
	}, opts...)

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
