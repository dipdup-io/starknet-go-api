package api

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

// GetBlockWithTxs -
func (api API) GetBlockWithTxs(ctx context.Context, bloxk BlockFilter, opts ...RequestOption) (*Response[BlockWithTxs], error) {
	if err := bloxk.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getBlockWithTxs", []any{bloxk}, opts...)

	var response Response[BlockWithTxs]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
