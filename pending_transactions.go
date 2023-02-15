package starknetgoapi

import "context"

// PendingTransactions - Returns the transactions in the transaction pool, recognized by this sequencer
func (api API) PendingTransactions(ctx context.Context, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_pendingTransactions", []any{}, opts...)

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
