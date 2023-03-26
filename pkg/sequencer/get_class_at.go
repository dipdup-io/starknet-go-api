package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetClassAt -
func (api API) GetClassAt(ctx context.Context, block data.BlockID, contractAddress string) (response data.Class, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := make(map[string]string)
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}
	args["contractAddress"] = contractAddress

	err = api.getFromFeederGateway(ctx, "get_full_contract", "", args, &response)
	return
}
