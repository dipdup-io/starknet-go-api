package starknetgoapi

import "context"

// StateUpdate -
type StateUpdate struct {
	BlockHash string    `json:"block_hash"`
	NewRoot   string    `json:"new_root"`
	OldRoot   string    `json:"old_root"`
	StateDiff StateDiff `json:"state_diff"`
}

// GetStateUpdate -
func (api API) GetStateUpdate(ctx context.Context, block BlockFilter, opts ...RequestOption) (*Response[StateUpdate], error) {
	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{block}, opts...)

	var response Response[StateUpdate]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
