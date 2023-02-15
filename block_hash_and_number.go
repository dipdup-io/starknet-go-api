package starknetgoapi

import "context"

// BlockHashAndNumber -
type BlockHashAndNumber struct {
	Hash   string `json:"block_hash"`
	Number uint64 `json:"block_number"`
}

// BlockHashAndNumber - Get the most recent accepted block hash and number
func (api API) BlockHashAndNumber(ctx context.Context, opts ...RequestOption) (*Response[BlockHashAndNumber], error) {
	request := api.prepareRequest(ctx, "starknet_blockHashAndNumber", []any{}, opts...)

	var response Response[BlockHashAndNumber]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
