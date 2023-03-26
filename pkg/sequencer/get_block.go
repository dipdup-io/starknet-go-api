package sequencer

import (
	"context"
	"fmt"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// Block -
type Block struct {
	Timestamp        int64              `json:"timestamp"`
	BlockNumber      uint64             `json:"block_number"`
	Status           string             `json:"status"`
	BlockHash        string             `json:"block_hash"`
	ParentHash       string             `json:"parent_block_hash"`
	NewRoot          string             `json:"state_root"`
	GasPrice         string             `json:"gas_price"`
	SequencerAddress string             `json:"sequencer_address"`
	StarknetVersion  *string            `json:"starknet_version,omitempty"`
	Transactions     []data.Transaction `json:"transactions"`
	Receipts         []Receipt          `json:"transaction_receipts"`
}

// Receipt -
type Receipt struct {
	TransactionIndex   uint64             `json:"transaction_index"`
	TransactionHash    string             `json:"transaction_hash"`
	L2ToL1Messages     []data.Message     `json:"l2_to_l1_messages"`
	L1ToL2Message      data.Message       `json:"l1_to_l2_consumed_message"`
	Events             []data.Event       `json:"events"`
	ExecutionResources ExecutionResources `json:"execution_resources"`
	ActualFee          data.Felt          `json:"actual_fee"`
}

// GetBlock - Gets block
func (api API) GetBlock(ctx context.Context, block data.BlockID) (response Block, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := make(map[string]string)
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}

	var cacheFileName string
	switch {
	case block.Number != nil:
		cacheFileName = fmt.Sprintf("%d.json", *block.Number)
	case block.Hash != "":
		cacheFileName = block.Hash
	}

	err = api.getFromFeederGateway(ctx, "get_block", cacheFileName, args, &response)
	return
}
