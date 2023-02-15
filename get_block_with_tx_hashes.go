package starknetgoapi

import "context"

// BlockWithTxHashes -
type BlockWithTxHashes struct {
	Status           string   `json:"status"`
	BlockHash        string   `json:"block_hash"`
	ParentHash       string   `json:"parent_hash"`
	BlockNumber      uint64   `json:"block_number"`
	NewRoot          string   `json:"new_root"`
	Timestamp        int64    `json:"timestamp"`
	SequencerAddress string   `json:"sequencer_address"`
	Transactions     []string `json:"transactions"`
}

// GetBlockWithTxHashesByNumber -
func (api API) GetBlockWithTxHashesByNumber(ctx context.Context, blockNumber uint64, opts ...RequestOption) (*Response[BlockWithTxHashes], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxHashes", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
	}, opts...)

	var response Response[BlockWithTxHashes]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetBlockWithTxHashesByHash -
func (api API) GetBlockWithTxHashesByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[BlockWithTxHashes], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxHashes", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
	}, opts...)

	var response Response[BlockWithTxHashes]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetLatestBlockWithTxHashes -
func (api API) GetLatestBlockWithTxHashes(ctx context.Context, opts ...RequestOption) (*Response[BlockWithTxHashes], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxHashes", []any{
		latest,
	}, opts...)

	var response Response[BlockWithTxHashes]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetPendingBlockWithTxHashes -
func (api API) GetPendingBlockWithTxHashes(ctx context.Context, opts ...RequestOption) (*Response[BlockWithTxHashes], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxHashes", []any{
		pending,
	}, opts...)

	var response Response[BlockWithTxHashes]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
