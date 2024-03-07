package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// BlockWithReceipt -
type BlockWithReceipt struct {
	Status           string     `json:"status"`
	BlockHash        string     `json:"block_hash"`
	ParentHash       string     `json:"parent_hash"`
	BlockNumber      uint64     `json:"block_number"`
	NewRoot          string     `json:"new_root"`
	Timestamp        int64      `json:"timestamp"`
	SequencerAddress string     `json:"sequencer_address"`
	Version          *string    `json:"starknet_version,omitempty"`
	L1GasPrice       L1GasPrice `json:"l1_gas_price"`
	Transactions     []Tx       `json:"transactions"`
}

type Tx struct {
	Transaction data.Transaction `json:"transaction"`
	Receipt     Receipt          `json:"receipt"`
}

// GetBlockWithReceipts -
func (api API) GetBlockWithReceipts(ctx context.Context, block data.BlockID, opts ...RequestOption) (*Response[BlockWithReceipt], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest("starknet_getBlockWithReceipts", []any{block}, opts...)

	var response Response[BlockWithReceipt]
	err := post(ctx, api, *request, &response)
	return &response, err
}
