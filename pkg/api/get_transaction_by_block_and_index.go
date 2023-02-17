package api

import "context"

// GetTransactionByBlockNumberAndIndex - Get the details of the transaction given by the identified block and index in that block. If no transaction is found, null is returned.
func (api API) GetTransactionByBlockNumberAndIndex(ctx context.Context, block BlockFilter, index uint64, opts ...RequestOption) (*Response[Transaction], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_getTransactionByBlockIdAndIndex", []any{
		block, index,
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
