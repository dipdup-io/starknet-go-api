package starknetgoapi

import "context"

// GetTransactionByBlockNumberAndIndex - Get the details of the transaction given by the identified block and index in that block. If no transaction is found, null is returned.
func (api API) GetTransactionByBlockNumberAndIndex(ctx context.Context, block BlockFilter, index uint64, opts ...RequestOption) (*Response[Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByBlockIdAndIndex", []any{
		block, index,
	}, opts...)

	var response Response[Transaction]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
