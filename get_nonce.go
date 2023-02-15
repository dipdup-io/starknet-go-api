package starknetgoapi

import "context"

// GetNonceByBlockNumber -
func (api API) GetNonceByBlockNumber(ctx context.Context, contract string, blockNumber uint64, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
		contract,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetNonceByBlockHash -
func (api API) GetNonceByBlockHash(ctx context.Context, contract, hash string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
		contract,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetLatestNonce -
func (api API) GetLatestNonce(ctx context.Context, contract string, opts ...RequestOption) (*Response[string], error) {
	request := api.prepareRequest(ctx, "starknet_getStorageAt", []any{
		latest,
		contract,
	}, opts...)

	var response Response[string]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
