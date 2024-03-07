package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// Receipt -
type Receipt struct {
	Type            string         `json:"type"`
	TransactionHash string         `json:"transaction_hash"`
	ActualFee       Fee            `json:"actual_fee"`
	ExecutionStatus string         `json:"execution_status"`
	FinalityStatus  string         `json:"finality_status"`
	BlockHash       string         `json:"block_hash"`
	BlockNumber     uint64         `json:"block_number"`
	MessagesSent    []data.Message `json:"messages_sent"`
	Events          []data.Event   `json:"events"`
	ContractAddress string         `json:"contract_address"`
}

type Fee struct {
	Amount data.Felt `json:"amount"`
	Unit   string    `json:"unit"`
}

// GetTransactionReceipts -
func (api API) GetTransactionReceipts(ctx context.Context, hash string, opts ...RequestOption) (*Response[Receipt], error) {
	request := api.prepareRequest("starknet_getTransactionReceipt", []any{
		hash,
	}, opts...)

	var response Response[Receipt]
	err := post(ctx, api, *request, &response)
	return &response, err
}
