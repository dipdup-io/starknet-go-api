# starknet-go-api
Go implementation of Starknet API

## Install

```bash
go get github.com/dipdup-io/starknet-api-go/pkg/api
```

## Usage

You can find example of usage [in our repository](/example/api/main.go). 

Library implements following methods:

```go
BlockHashAndNumber(ctx context.Context, opts ...RequestOption) (*Response[BlockHashAndNumber], error)
BlockNumber(ctx context.Context, opts ...RequestOption) (*Response[uint64], error)
Call(ctx context.Context, params CallRequest, block BlockFilter, opts ...RequestOption) (*Response[[]string], error)
ChainID(ctx context.Context, opts ...RequestOption) (*Response[string], error)
EstimateFee(ctx context.Context, tx Transaction, block BlockFilter, opts ...RequestOption) (*Response[EstmatedGas], error)
GetBlockTransactionCount(ctx context.Context, block BlockFilter, opts ...RequestOption) (*Response[uint64], error)
GetBlockWithTxHashes(ctx context.Context, block BlockFilter, opts ...RequestOption) (*Response[BlockWithTxHashes], error)
GetBlockWithTxs(ctx context.Context, bloxk BlockFilter, opts ...RequestOption) (*Response[BlockWithTxs], error)
GetClassAt(ctx context.Context, block BlockFilter, contractAddress string, opts ...RequestOption) (*Response[Class], error)
GetClassHashAt(ctx context.Context, block BlockFilter, contractAddress string, opts ...RequestOption) (*Response[string], error)
GetClass(ctx context.Context, block BlockFilter, classHash string, opts ...RequestOption) (*Response[Class], error)
GetEvents(ctx context.Context, filters EventsFilters, opts ...RequestOption) (*Response[EventsResponse], error) 
GetNonce(ctx context.Context, contract string, block BlockFilter, opts ...RequestOption) (*Response[string], error)
GetStateUpdate(ctx context.Context, block BlockFilter, opts ...RequestOption) (*Response[StateUpdate], error)
GetStorageAt(ctx context.Context, contract, key string, block BlockFilter, opts ...RequestOption) (*Response[string], error)
GetTransactionByBlockNumberAndIndex(ctx context.Context, block BlockFilter, index uint64, opts ...RequestOption) (*Response[Transaction], error)
GetTransactionByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[Transaction], error)
GetTransactionReceipts(ctx context.Context, hash string, opts ...RequestOption) (*Response[Receipt], error)
PendingTransactions(ctx context.Context, opts ...RequestOption) (*Response[Transaction], error)
Syncing(ctx context.Context, opts ...RequestOption) (*Response[Syncing], error)
```

The most usable filter is `BlockFilter`. It has the following definition:

```go
// BlockFilter -
type BlockFilter struct {
	Hash   string
	Number uint64
	String string
}
```

You can set `Hash` or `Number` of block or set `String` to `latest` or `pending` value to filter by them. But you should set only one field of `BlockFilter` structure. If you set 2 fields you'll receive validation error. For example:

```go
// for hash
api.GetBlockWithTxs(ctx, BlockFilter{Hash: "some_hash"})
// for number
api.GetBlockWithTxs(ctx, BlockFilter{Number: 100})
// for latest
api.GetBlockWithTxs(ctx, BlockFilter{String: Latest})
// for pending
api.GetBlockWithTxs(ctx, BlockFilter{String: Pending})

// wrong call
api.GetBlockWithTxs(ctx, BlockFilter{Hash: "some_hash", Number: 100})
```

Simple usage example:

```go
package main

import (
	"context"
	"log"

	starknet "github.com/dipdup-io/starknet-go-api/pkg/api"
)

func main() {
	api := starknet.NewAPI("HERE_LINK_TO_API")
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
```