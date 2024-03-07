package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// PendingTransactions - Returns the transactions in the transaction pool, recognized by this sequencer
func (api API) PendingTransactions(ctx context.Context, opts ...RequestOption) (*Response[data.Transaction], error) {
	request := api.prepareRequest("starknet_PendingTransactions", []any{}, opts...)

	var response Response[data.Transaction]
	err := post(ctx, api, *request, &response)
	return &response, err
}
