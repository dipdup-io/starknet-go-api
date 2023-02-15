package starknetgoapi

import "context"

// BlockWithTxs -
type BlockWithTxs struct {
	Status           string        `json:"status"`
	BlockHash        string        `json:"block_hash"`
	ParentHash       string        `json:"parent_hash"`
	BlockNumber      uint64        `json:"block_number"`
	NewRoot          string        `json:"new_root"`
	Timestamp        int64         `json:"timestamp"`
	SequencerAddress string        `json:"sequencer_address"`
	Transactions     []Transaction `json:"transactions"`
}

// GetBlockWithTxsByNumber -
func (api API) GetBlockWithTxsByNumber(ctx context.Context, blockNumber uint64, opts ...RequestOption) (*Response[BlockWithTxs], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxs", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
	}, opts...)

	var response Response[BlockWithTxs]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetBlockWithTxsByHash -
func (api API) GetBlockWithTxsByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[BlockWithTxs], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxs", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
	}, opts...)

	var response Response[BlockWithTxs]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetLatestBlockWithTxs -
func (api API) GetLatestBlockWithTxs(ctx context.Context, opts ...RequestOption) (*Response[BlockWithTxs], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxs", []any{
		latest,
	}, opts...)

	var response Response[BlockWithTxs]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetPendingBlockWithTxs -
func (api API) GetPendingBlockWithTxs(ctx context.Context, opts ...RequestOption) (*Response[BlockWithTxs], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxs", []any{
		pending,
	}, opts...)

	var response Response[BlockWithTxs]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
