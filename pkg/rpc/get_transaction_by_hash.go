package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetTransactionByHash -
func (api API) GetTransactionByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[data.Transaction], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionByHash", []any{
		hash,
	}, opts...)

	var response Response[data.Transaction]
	err := post(ctx, api, *request, &response)
	return &response, err
}
