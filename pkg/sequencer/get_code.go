package sequencer

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/abi"
	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// Code -
type Code struct {
	ByteCode []string `json:"bytecode"`
	Abi      abi.Abi  `json:"abi"`
}

// GetBlock - Gets code of contract
func (api API) GetCode(ctx context.Context, block data.BlockID, contractAddress string) (response Code, err error) {
	if err := block.Validate(); err != nil {
		return response, err
	}

	args := map[string]string{
		"contractAddress": contractAddress,
	}
	if name, value := block.GetArg(); name != "" {
		args[name] = value
	}

	err = api.getFromFeederGateway(ctx, "get_code", "", args, &response)
	return
}
