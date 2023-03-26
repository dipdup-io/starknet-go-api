package sequencer

import (
	"context"
)

// TraceTransaction - Gets the transaction trace from hash
func (api API) TraceTransaction(ctx context.Context, hash string) (response Trace, err error) {
	args := map[string]string{
		"transactionHash": hash,
	}
	err = api.getFromFeederGateway(ctx, "get_transaction_trace", "", args, &response)
	return
}
