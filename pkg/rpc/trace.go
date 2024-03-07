package api

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

type Trace struct {
	TraceRoot       TraceRoot `json:"trace_root"`
	TransactionHash data.Felt `json:"transaction_hash"`
}

type TraceRoot struct {
	Type                  string `json:"type"`
	ExecuteInvocation     *Call  `json:"execute_invocation,omitempty"`
	ConstructorInvocation *Call  `json:"constructor_invocation,omitempty"`
	ValidateInvocation    *Call  `json:"validate_invocation,omitempty"`
	FeeTransferInvocation *Call  `json:"fee_transfer_invocation,omitempty"`
}

type Call struct {
	CallerAddress      data.Felt          `json:"caller_address"`
	ContractAddress    data.Felt          `json:"contract_address"`
	CallType           string             `json:"call_type"`
	ClassHash          data.Felt          `json:"class_hash"`
	EntryPointSelector data.Felt          `json:"entry_point_selector"`
	EntryPointType     string             `json:"entry_point_type"`
	RevertReason       string             `json:"revert_reason,omitempty"`
	Calldata           []data.Felt        `json:"calldata"`
	Result             []data.Felt        `json:"result"`
	Calls              []Call             `json:"calls"`
	Events             []data.Event       `json:"events"`
	Messages           []data.Message     `json:"messages"`
	ExecutionResources ExecutionResources `json:"execution_resources"`
}

type ExecutionResources struct {
	Steps                         int `json:"steps"`
	MemoryHoles                   int `json:"memory_holes"`
	PedersenBuiltinApplications   int `json:"pedersen_builtin_applications"`
	RangeCheckBuiltinApplications int `json:"range_check_builtin_applications"`
}

// Trace -
func (api API) Trace(ctx context.Context, block data.BlockID, opts ...RequestOption) (*Response[[]Trace], error) {
	if err := block.Validate(); err != nil {
		return nil, err
	}

	request := api.prepareRequest("starknet_traceBlockTransactions", []any{block}, opts...)

	var response Response[[]Trace]
	err := post(ctx, api, *request, &response)
	return &response, err
}
