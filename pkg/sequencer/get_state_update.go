package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetStateUpdate -
func (api API) GetStateUpdate(ctx context.Context, block data.BlockID) (response data.StateUpdate, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := map[string]string{}
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}

	err = api.getFromFeederGateway(ctx, "get_state_update", args, &response)
	return
}
