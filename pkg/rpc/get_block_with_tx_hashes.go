package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

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
func (api API) GetBlockWithTxHashes(ctx context.Context, block data.BlockID, opts ...RequestOption) (*Response[BlockWithTxHashes], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getBlockWithTxHashes", []any{block}, opts...)

	var response Response[BlockWithTxHashes]
	err := post(ctx, api, *request, &response)
	return &response, err
}
