package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// EventsResponse -
type EventsResponse struct {
	Events            []data.Event `json:"events"`
	ContinuationToken string       `json:"continuation_token"`
}

// EventsFilters -
type EventsFilters struct {
	FromBlock *data.BlockID `json:"from_block,omitempty"`
	ToBlock   *data.BlockID `json:"to_block,omitempty"`
	Address   string        `json:"address,omitempty"`
	Keys      [][]string    `json:"keys"`
}

// GetEvents - Returns all event objects matching the conditions in the provided filter
func (api API) GetEvents(ctx context.Context, filters EventsFilters, opts ...RequestOption) (*Response[EventsResponse], error) {
	if filters.FromBlock != nil {
		if err := filters.FromBlock.Validate(); err != nil {
			return nil, err
		}
	}
	if filters.ToBlock != nil {
		if err := filters.ToBlock.Validate(); err != nil {
			return nil, err
		}
	}

	request := api.prepareRequest(ctx, "", []any{filters}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[EventsResponse]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
