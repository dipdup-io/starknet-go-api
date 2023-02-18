package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// GetNonce - Gets nonce for address
func (api API) GetNonce(ctx context.Context, block data.BlockID, contractAddress string) (response string, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := map[string]string{
		"contractAddress": contractAddress,
	}
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}

	err = api.getFromFeederGateway(ctx, "get_nonce", args, &response)
	return
}
