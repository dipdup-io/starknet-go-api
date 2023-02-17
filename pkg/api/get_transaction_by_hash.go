package api

import "context"

// GetTransactionByHash -
func (api API) GetTransactionByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByHash", []any{
		hash,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
