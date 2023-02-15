package starknetgoapi

import "context"

// Syncing -
type Syncing struct {
	StartingBlockNum  string `json:"starting_block_num"`
	CurrentBlockNum   string `json:"current_block_num"`
	HighestBlockNum   string `json:"highest_block_num"`
	StartingBlockHash string `json:"starting_block_hash"`
	CurrentBlockHash  string `json:"current_block_hash"`
	HighestBlockHash  string `json:"highest_block_hash"`
}

// TODO: handle false
// Syncing - Returns an object about the sync status, or false if the node is not synching
func (api API) Syncing(ctx context.Context, opts ...RequestOption) (*Response[Syncing], error) {
	request := api.prepareRequest(ctx, "starknet_syncing", []any{}, opts...)

	var response Response[Syncing]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
