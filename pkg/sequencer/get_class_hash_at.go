package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetClassHashAt -
func (api API) GetClassHashAt(ctx context.Context, block data.BlockID, contractAddress string) (response string, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := make(map[string]string)
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}
	args["contractAddress"] = contractAddress

	err = api.getFromFeederGateway(ctx, "get_class_hash_at", args, &response)
	return
}
