package main

import (
	"context"
	"log"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
	starknet "github.com/dipdup-io/starknet-go-api/pkg/rpc"
)

func main() {
	api := starknet.NewAPI("LINK_TO_NODE_RPC")
	ctx := context.Background()

	blockNumber := uint64(100)

	response, err := api.GetBlockTransactionCount(ctx, data.BlockID{
		Number: blockNumber,
	}, starknet.WithTimeout(10))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("transaction count in block %d = %d", blockNumber, response.Result)
}
