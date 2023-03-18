package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// CallContract -
func (api API) CallContract(ctx context.Context, block data.BlockID, contractAddress, entrypointSelector string, calldata []string) (response Response[[]data.Felt], err error) {
	if err = block.Validate(); err != nil {
		return
	}

	args := make(map[string]string)
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}

	body := map[string]any{
		"signature":            []string{},
		"contract_address":     contractAddress,
		"entry_point_selector": entrypointSelector,
		"calldata":             calldata,
	}

	err = api.postToFeederGateway(ctx, "call_contract", args, body, &response)
	return
}
