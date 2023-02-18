# Sequencer API

Package implements wrapper over Sequencer API. Package is under development.

## Install

```bash
go get github.com/dipdup-io/starknet-api-go/pkg/sequencer
```

## Usage

Full example of usage can be found [here](/example/sequencer/main.go).


### Methods

Wrapper implements next methods:

```go

CallContract(ctx context.Context, block data.BlockID, contractAddress, entrypointSelector string, calldata []string) (response Response[[]string], err error)

GetTransaction(ctx context.Context, hash string) (response Trace, err error) 

GetCode(ctx context.Context, block data.BlockID, contractAddress string) (response Code, err error)

GetNonce(ctx context.Context, block data.BlockID, contractAddress string) (response string, err error)

TraceBlock(ctx context.Context, block data.BlockID) (response TraceResponse, err error)

TraceTransaction(ctx context.Context, hash string) (response Trace, err error)
```

### Creation

First of all you should import package in your code:

```go
starknetData "github.com/dipdup-io/starknet-go-api/pkg/data"
sequencerAPI "github.com/dipdup-io/starknet-go-api/pkg/sequencer"
```

Then create `API` object:

```go
api := sequencerAPI.NewAPI("https://alpha-mainnet.starknet.io/gateway", "https://alpha-mainnet.starknet.io/feeder_gateway")
```

And call any method:

```go
response, err := api.CallContract(ctx, starknetData.BlockID{
    String: starknetData.Latest,
},
    "0x233084545b87df4940643bdcc5ff959f8371d2f388ae5f05c8c19eea7059c1a",
    "0x26813d396fdb198e9ead934e4f7a592a8b88a059e45ab0eb6ee53494e8d45b0",
    []string{
        "5",
    })
if err != nil {
    log.Panic(err)
}
```
