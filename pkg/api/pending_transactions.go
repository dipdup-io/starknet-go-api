package api

import "context"

// PendingTransactions - Returns the transactions in the transaction pool, recognized by this sequencer
func (api API) PendingTransactions(ctx context.Context, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_PendingTransactions", []any{}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
