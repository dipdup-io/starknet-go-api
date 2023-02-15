package starknetgoapi

import "context"

// GetStorageByNumber -
func (api API) GetStorageByNumber(ctx context.Context, contract, key string, blockNumber uint64, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		contract,
		key,
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetStorageByHash -
func (api API) GetStorageByHash(ctx context.Context, contract, key, hash string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		contract,
		key,
		&BlockRequest{
			BlockHash: &hash,
		},
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetLatestStorage -
func (api API) GetLatestStorage(ctx context.Context, contract, key string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		contract,
		key,
		latest,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetPendingStorage -
func (api API) GetPendingStorage(ctx context.Context, contract, key string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		contract,
		key,
		pending,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
