package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// Receipt -
type Receipt struct {
	Type            string         `json:"type"`
	TransactionHash string         `json:"transaction_hash"`
	ActualFee       string         `json:"actual_fee"`
	Status          string         `json:"status"`
	BlockHash       string         `json:"block_hash"`
	BlockNumber     uint64         `json:"block_number"`
	MessagesSent    []data.Message `json:"messages_sent"`
	Events          []data.Event   `json:"events"`
	ContractAddress string         `json:"contract_address"`
}

// GetTransactionReceipts -
func (api API) GetTransactionReceipts(ctx context.Context, hash string, opts ...RequestOption) (*Response[Receipt], error) {
	request := api.prepareRequest(ctx, "starknet_getTransactionReceipt", []any{
		hash,
	}, opts...)

	if api.rateLimit != nil {
		if err := api.rateLimit.Wait(ctx); err != nil {
			return nil, err
		}
	}

	var response Response[Receipt]
	err := post(ctx, api.client, api.baseURL, *request, &response)
	return &response, err
}
