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

// GetBlockWithTxHashes -
func (api API) GetBlockWithTxHashes(ctx context.Context, filters BlockFilter, opts ...RequestOption) (*Response[BlockWithTxHashes], error) {
	request := api.prepareRequest(ctx, "starknet_getBlockWithTxHashes", []any{filters}, opts...)

	var response Response[BlockWithTxHashes]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
