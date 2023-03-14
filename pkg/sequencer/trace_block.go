package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// TraceResponse -
type TraceResponse struct {
	Traces []Trace `json:"traces"`
}

// Invocation -
type Invocation struct {
	CallerAddress      data.Felt          `json:"caller_address"`
	ContractAddress    data.Felt          `json:"contract_address"`
	Calldata           []string           `json:"calldata"`
	CallType           string             `json:"call_type"`
	ClassHash          data.Felt          `json:"class_hash"`
	Selector           data.Felt          `json:"selector"`
	EntrypointType     string             `json:"entry_point_type"`
	Result             []string           `json:"result"`
	ExecutionResources ExecutionResources `json:"execution_resources"`
	InternalCalls      []Invocation       `json:"internal_calls"`
	Events             []data.Event       `json:"events"`
	Messages           []data.Message     `json:"messages"`
}

// ExecutionResources -
type ExecutionResources struct {
	NSteps                 int                    `json:"n_steps"`
	BuiltinInstanceCounter BuiltinInstanceCounter `json:"builtin_instance_counter"`
	NMemoryHoles           int                    `json:"n_memory_holes"`
}

// BuiltinInstanceCounter -
type BuiltinInstanceCounter struct {
	RangeCheckBuiltin int `json:"range_check_builtin"`
	EcdsaBuiltin      int `json:"ecdsa_builtin"`
	PedersenBuiltin   int `json:"pedersen_builtin"`
	BitwiseBuiltin    int `json:"bitwise_builtin"`
}

// Trace -
type Trace struct {
	ValidateInvocation    *Invocation `json:"validate_invocation,omitempty"`
	FunctionInvocation    *Invocation `json:"function_invocation,omitempty"`
	FeeTransferInvocation *Invocation `json:"fee_transfer_invocation,omitempty"`
	Signature             []string    `json:"signature"`
	TransactionHash       data.Felt   `json:"transaction_hash"`
}

// TraceBlock -
func (api API) TraceBlock(ctx context.Context, block data.BlockID) (response TraceResponse, err error) {
	if err = block.Validate(); err != nil {
		return
	}

	args := make(map[string]string)
	if blockArgName, blockArgValue := block.GetArg(); blockArgName != "" {
		args[blockArgName] = blockArgValue
	}
	err = api.getFromFeederGateway(ctx, "get_block_traces", args, &response)

	return
}
