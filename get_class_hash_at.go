package starknetgoapi

import "context"

// GetClassHashByBlockNumber -
func (api API) GetClassHashByBlockNumber(ctx context.Context, blockNumber uint64, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
		contractAddress,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassHashByBlockHash -
func (api API) GetClassHashByBlockHash(ctx context.Context, hash, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
		contractAddress,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassHashByLatestBlock -
func (api API) GetClassHashByLatestBlock(ctx context.Context, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		latest,
		contractAddress,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassHashByPendingBlock -
func (api API) GetClassHashByPendingBlock(ctx context.Context, contractAddress string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getClassHashAt", []any{
		pending,
		contractAddress,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
