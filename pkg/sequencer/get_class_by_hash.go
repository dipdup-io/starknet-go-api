package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetClassByHash -
func (api API) GetClassByHash(ctx context.Context, block data.BlockID, classHash string) (response data.Class, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := make(map[string]string)
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}
	args["classHash"] = classHash

	err = api.getFromFeederGateway(ctx, "get_class_by_hash", args, &response)
	return
}
