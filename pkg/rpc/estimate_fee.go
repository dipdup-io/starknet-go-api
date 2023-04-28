package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// EstmatedGas -
type EstmatedGas struct {
	GasConsumed string `json:"gas_consumed"`
	GasPrice    string `json:"gas_price"`
	OverallFee  string `json:"overall_fee"`
}

// EstimateFee - estimates the resources required by a transaction relative to a given state
func (api API) EstimateFee(ctx context.Context, tx data.Transaction, block data.BlockID, opts ...RequestOption) (*Response[EstmatedGas], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_estimateFee", []any{tx})

	var response Response[EstmatedGas]
	err := post(ctx, api, *request, &response)
	return &response, err
}
