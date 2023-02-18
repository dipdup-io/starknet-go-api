package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// InvokeV0 -
func (api API) InvokeV0(ctx context.Context, tx data.InvokeV0) (string, error) {
	body := map[string]any{
		"type":             data.TransactionTypeInvokeFunction,
		"contract_address": tx.ContractAddress,
		"calldata":         tx.Calldata,
		"signature":        tx.Signature,
		"nonce":            tx.Nonce,
		"max_fee":          tx.MaxFee,
		"version":          data.Version0,
	}

	var hash string
	err := api.postToGateway(ctx, "add_transaction", nil, body, &hash)
	return hash, err
}
