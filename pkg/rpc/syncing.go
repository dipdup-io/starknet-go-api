package api

import (
	"context"

	"github.com/goccy/go-json"
)

// Syncing -
type Syncing struct {
	Synced            bool   `json:"-"`
	StartingBlockNum  string `json:"starting_block_num"`
	CurrentBlockNum   string `json:"current_block_num"`
	HighestBlockNum   string `json:"highest_block_num"`
	StartingBlockHash string `json:"starting_block_hash"`
	CurrentBlockHash  string `json:"current_block_hash"`
	HighestBlockHash  string `json:"highest_block_hash"`
}

// UnmarshalJSON -
func (s *Syncing) UnmarshalJSON(data []byte) error {
	type buf Syncing
	if err := json.Unmarshal(data, (*buf)(s)); err == nil {
		s.Synced = true
		return nil
	}

	var synced bool
	if err := json.Unmarshal(data, &synced); err != nil {
		return err
	}

	s.Synced = synced
	return nil
}

// Syncing - Returns an object about the sync status, or false if the node is not synching
func (api API) Syncing(ctx context.Context, opts ...RequestOption) (*Response[Syncing], error) {
	request := api.prepareRequest(ctx, "starknet_syncing", []any{}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[Syncing]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
