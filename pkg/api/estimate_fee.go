package api

import "context"

// EstmatedGas -
type EstmatedGas struct {
	GasConsumed string `json:"gas_consumed"`
	GasPrice    string `json:"gas_price"`
	OverallFee  string `json:"overall_fee"`
}

// EstimateFee - estimates the resources required by a transaction relative to a given state
func (api API) EstimateFee(ctx context.Context, tx Transaction, block BlockFilter, opts ...RequestOption) (*Response[EstmatedGas], error) {
	if err := block.validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest(ctx, "starknet_estimateFee", []any{tx})

	var response Response[EstmatedGas]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
