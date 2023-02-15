package starknetgoapi

import "context"

// StateUpdate -
type StateUpdate struct {
	BlockHash string    `json:"block_hash"`
	NewRoot   string    `json:"new_root"`
	OldRoot   string    `json:"old_root"`
	StateDiff StateDiff `json:"state_diff"`
}

// GetStateUpdateByNumber -
func (api API) GetStateUpdateByNumber(ctx context.Context, blockNumber uint64, opts ...RequestOption) (*Response[StateUpdate], error) {
	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{
		&BlockRequest{
			BlockNumber: &blockNumber,
		},
	}, opts...)

	var response Response[StateUpdate]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetStateUpdateByHash -
func (api API) GetStateUpdateByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[StateUpdate], error) {
	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{
		&BlockRequest{
			BlockHash: &hash,
		},
	}, opts...)

	var response Response[StateUpdate]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetLatestStateUpdate -
func (api API) GetLatestStateUpdate(ctx context.Context, opts ...RequestOption) (*Response[StateUpdate], error) {
	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{
		latest,
	}, opts...)

	var response Response[StateUpdate]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}

// GetPendingStateUpdate -
func (api API) GetPendingStateUpdate(ctx context.Context, opts ...RequestOption) (*Response[StateUpdate], error) {
	request := api.prepareRequest(ctx, "starknet_getStateUpdate", []any{
		pending,
	}, opts...)

	var response Response[StateUpdate]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
