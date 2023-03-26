package sequencer

import (
	"context"
	"fmt"

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

	var cacheFileName string
	switch {
	case block.Number != nil:
		cacheFileName = fmt.Sprintf("%s_%d.json", classHash, *block.Number)
	case block.Hash != "":
		cacheFileName = fmt.Sprintf("%s_%s.json", classHash, block.Hash)
	case block.String != "":
		cacheFileName = fmt.Sprintf("%s_%s.json", classHash, block.String)
	}

	err = api.getFromFeederGateway(ctx, "get_class_by_hash", cacheFileName, args, &response)
	return
}
