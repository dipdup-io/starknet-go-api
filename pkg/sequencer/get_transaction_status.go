package sequencer

import (
	"context"
)

// TransactionStatus -
type TransactionStatus struct {
	Status    string `json:"tx_status"`
	BlockHash string `json:"block_hash"`
}

// GetTransactionStatus - Gets nonce for address
func (api API) GetTransactionStatus(ctx context.Context, txHash string) (response TransactionStatus, err error) {
	args := map[string]string{
		"transactionHash": txHash,
	}

	err = api.getFromFeederGateway(ctx, "get_transaction_status", "", args, &response)
	return
}
