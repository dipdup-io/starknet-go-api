# Data

Package contains general structures.

## BlockID

The most usable filter is `BlockID`. It has the following definition:

```go
// BlockID -
type BlockID struct {
	Hash   string
	Number uint64
	String string
}
```

You can set `Hash` or `Number` of block or set `String` to `latest` or `pending` value to filter by them. But you should set only one field of `BlockID` structure. If you set 2 fields you'll receive validation error. For example:

```go
// for hash
api.GetBlockWithTxs(ctx, BlockID{Hash: "some_hash"})
// for number
api.GetBlockWithTxs(ctx, BlockID{Number: 100})
// for latest
api.GetBlockWithTxs(ctx, BlockID{String: Latest})
// for pending
api.GetBlockWithTxs(ctx, BlockID{String: Pending})

// wrong call
api.GetBlockWithTxs(ctx, BlockID{Hash: "some_hash", Number: 100})
```
