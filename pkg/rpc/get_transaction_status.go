package api

import "context"

type TransactionStatus struct {
	Finality  string `json:"finality_status"`
	Execution string `json:"execution_status"`
}

// GetTransactionStatus -
func (api API) GetTransactionStatus(ctx context.Context, hash string, opts ...RequestOption) (*Response[TransactionStatus], error) {
	request := api.prepareRequest("starknet_getTransactionReceipt", []any{
		hash,
	}, opts...)

	var response Response[TransactionStatus]
	err := post(ctx, api, *request, &response)
	return &response, err
}
