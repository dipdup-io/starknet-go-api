package main

import (
	"context"
	"log"

	starknet "github.com/dipdup-io/starknet-go-api"
)

func main() {
	api := starknet.NewAPI("https://starknet-testnet.public.blastapi.io")
	ctx := context.Background()

	blockNumber := uint64(100)

	response, err := api.GetBlockTransactionCount(ctx, starknet.BlockFilter{
		Number: blockNumber,
	}, starknet.WithTimeout(10))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("transaction count in block %d = %d", blockNumber, response.Result)
}
