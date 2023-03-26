package sequencer

import (
	"context"
	"fmt"

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

	var cacheFileName string
	switch {
	case block.Number != nil:
		cacheFileName = fmt.Sprintf("%d.json", *block.Number)
	case block.Hash != "":
		cacheFileName = block.Hash
	}

	err = api.getFromFeederGateway(ctx, "get_state_update", cacheFileName, args, &response)
	return
}
