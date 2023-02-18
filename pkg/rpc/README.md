# Node RPC API

Go implementation of Starknet node API

## Install

```bash
go get github.com/dipdup-io/starknet-api-go/pkg/rpc
```

## Usage

You can find example of usage [in our repository](/example/rpc/main.go). 


### Methods
Library implements following methods:

```go
BlockHashAndNumber(ctx context.Context, opts ...RequestOption) (*Response[BlockHashAndNumber], error)

BlockNumber(ctx context.Context, opts ...RequestOption) (*Response[uint64], error)

Call(ctx context.Context, params CallRequest, block BlockID, opts ...RequestOption) (*Response[[]string], error)

ChainID(ctx context.Context, opts ...RequestOption) (*Response[string], error)

EstimateFee(ctx context.Context, tx Transaction, block BlockID, opts ...RequestOption) (*Response[EstmatedGas], error)

GetBlockTransactionCount(ctx context.Context, block BlockID, opts ...RequestOption) (*Response[uint64], error)

GetBlockWithTxHashes(ctx context.Context, block BlockID, opts ...RequestOption) (*Response[BlockWithTxHashes], error)

GetBlockWithTxs(ctx context.Context, bloxk BlockID, opts ...RequestOption) (*Response[BlockWithTxs], error)

GetClassAt(ctx context.Context, block BlockID, contractAddress string, opts ...RequestOption) (*Response[Class], error)

GetClassHashAt(ctx context.Context, block BlockID, contractAddress string, opts ...RequestOption) (*Response[string], error)

GetClass(ctx context.Context, block BlockID, classHash string, opts ...RequestOption) (*Response[Class], error)

GetEvents(ctx context.Context, filters EventsFilters, opts ...RequestOption) (*Response[EventsResponse], error) 

GetNonce(ctx context.Context, contract string, block BlockID, opts ...RequestOption) (*Response[string], error)

GetStateUpdate(ctx context.Context, block BlockID, opts ...RequestOption) (*Response[StateUpdate], error)

GetStorageAt(ctx context.Context, contract, key string, block BlockID, opts ...RequestOption) (*Response[string], error)

GetTransactionByBlockNumberAndIndex(ctx context.Context, block BlockID, index uint64, opts ...RequestOption) (*Response[Transaction], error)

GetTransactionByHash(ctx context.Context, hash string, opts ...RequestOption) (*Response[Transaction], error)

GetTransactionReceipts(ctx context.Context, hash string, opts ...RequestOption) (*Response[Receipt], error)

PendingTransactions(ctx context.Context, opts ...RequestOption) (*Response[Transaction], error)

Syncing(ctx context.Context, opts ...RequestOption) (*Response[Syncing], error)
```

### Creation

First of all you should import package in your code:


```go
starknetData "github.com/dipdup-io/starknet-go-api/pkg/data"
rpc "github.com/dipdup-io/starknet-go-api/pkg/rpc"
```

Then create `API` object:

```go
api := rpc.NewAPI("LINK_TO_NODE_RPC")
```

And call any method:

```go
response, err := api.GetBlockTransactionCount(ctx, starknetData.BlockID{
    Number: 100,
}, rpc.WithTimeout(10))
if err != nil {
    log.Panic(err)
}
```

### Timeout

If you need timeout on your request you can use option `WithTimeout` and set timeout by passing context:

```go
// by option
response, err := api.GetBlockTransactionCount(ctx, starknetData.BlockID{
    Number: 100,
}, rpc.WithTimeout(10))
if err != nil {
    log.Panic(err)
}

// by context
requestCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
defer cancel()

response, err := api.GetBlockTransactionCount(requestCtx, starknetData.BlockID{
    Number: 100,
})
if err != nil {
    log.Panic(err)
}
```

If you pass both variants of timeout the least will be used.