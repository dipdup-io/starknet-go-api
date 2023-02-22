package main

import (
	"context"
	"log"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
	starknet "github.com/dipdup-io/starknet-go-api/pkg/sequencer"
)

func main() {
	api := starknet.NewAPI("https://alpha-mainnet.starknet.io/gateway", "https://alpha-mainnet.starknet.io/feeder_gateway")
	ctx := context.Background()

	response, err := api.CallContract(ctx, data.BlockID{
		String: data.Latest,
	},
		"0x233084545b87df4940643bdcc5ff959f8371d2f388ae5f05c8c19eea7059c1a",
		"0x26813d396fdb198e9ead934e4f7a592a8b88a059e45ab0eb6ee53494e8d45b0",
		[]string{
			"5",
		})
	if err != nil {
		log.Panic(err)
	}

	log.Printf("call contract result %##v", response.Result)

	number := uint64(1)
	block, err := api.GetBlock(ctx, data.BlockID{
		Number: &number,
	})
	if err != nil {
		log.Panic(err)
	}

	log.Printf("block received: %##v", block)
}
