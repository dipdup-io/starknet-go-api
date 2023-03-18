package presets

import (
	"github.com/dipdup-io/starknet-go-api/pkg/data"
)

// CallOptions -
type CallOptions struct {
	block data.BlockID
}

// NewCallOptions -
func NewCallOptions(opts ...CallOption) CallOptions {
	options := CallOptions{
		block: data.BlockID{
			String: data.Latest,
		},
	}

	for i := range opts {
		opts[i](&options)
	}

	return options
}

// CallOption -
type CallOption func(*CallOptions)

// WithBlockID -
func WithBlockID(block data.BlockID) CallOption {
	return func(options *CallOptions) {
		options.block = block
	}
}
