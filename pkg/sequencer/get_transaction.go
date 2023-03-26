package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetTransactionResponse -
type GetTransactionResponse struct {
	Status           string           `json:"status"`
	BlockHash        string           `json:"block_hash"`
	BlockNumber      uint64           `json:"block_number"`
	TransactionIndex uint64           `json:"transaction_index"`
	Transaction      data.Transaction `json:"transaction"`
}

// GetTransaction - Gets the transaction by hash
func (api API) GetTransaction(ctx context.Context, hash string) (response Trace, err error) {
	args := map[string]string{
		"transactionHash": hash,
	}
	err = api.getFromFeederGateway(ctx, "get_transaction", "", args, &response)
	return
}
