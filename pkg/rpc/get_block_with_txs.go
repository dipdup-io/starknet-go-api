package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// BlockWithTxs -
type BlockWithTxs struct {
	Status           string             `json:"status"`
	BlockHash        string             `json:"block_hash"`
	ParentHash       string             `json:"parent_hash"`
	BlockNumber      uint64             `json:"block_number"`
	NewRoot          string             `json:"new_root"`
	Timestamp        int64              `json:"timestamp"`
	SequencerAddress string             `json:"sequencer_address"`
	Version          *string            `json:"starknet_version,omitempty"`
	L1GasPrice       L1GasPrice         `json:"l1_gas_price"`
	Transactions     []data.Transaction `json:"transactions"`
}

// GetBlockWithTxs -
func (api API) GetBlockWithTxs(ctx context.Context, block data.BlockID, opts ...RequestOption) (*Response[BlockWithTxs], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getBlockWithTxs", []any{block}, opts...)

	var response Response[BlockWithTxs]
	err := post(ctx, api, *request, &response)
	return &response, err
}
