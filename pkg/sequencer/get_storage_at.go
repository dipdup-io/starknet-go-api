package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetStorageAt - 'key' is decimal string
func (api API) GetStorageAt(ctx context.Context, block data.BlockID, contractAddress, key string) (response string, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := map[string]string{
		"contractAddress": contractAddress,
		"key":             key,
	}
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}

	err = api.getFromFeederGateway(ctx, "get_storage_at", args, &response)
	return
}
