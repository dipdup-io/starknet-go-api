package starknetgoapi

import "context"

// GetClassAtBlockNumber -
func (api API) GetClassAtBlockNumber(ctx context.Context, blockNumber uint64, contractAddress string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
		contractAddress,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassAtBlockHash -
func (api API) GetClassAtBlockHash(ctx context.Context, hash, contractAddress string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
		contractAddress,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassAtLatestBlock -
func (api API) GetClassAtLatestBlock(ctx context.Context, contractAddress string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		latest,
		contractAddress,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetClassAtPendingBlock -
func (api API) GetClassAtPendingBlock(ctx context.Context, contractAddress string, opts ...RequestOption) (*Response[Class], error) {
	request := api.prepareRequest(ctx, "starknet_getClassAt", []any{
		pending,
		contractAddress,
	}, opts...)

	var response Response[Class]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
