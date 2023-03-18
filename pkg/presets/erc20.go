package presets

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
	"github.com/dipdup-io/starknet-go-api/pkg/sequencer"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// ERC20 -
type ERC20 struct {
	api      sequencer.API
	contract data.Felt

	selectors map[string]string
}

// NewERC20 -
func NewERC20(api sequencer.API, contract data.Felt) ERC20 {
	return ERC20{
		api:      api,
		contract: contract,
		selectors: map[string]string{
			"balanceOf":   encoding.GetSelectorWithPrefixFromName("balanceOf"),
			"name":        encoding.GetSelectorWithPrefixFromName("name"),
			"symbol":      encoding.GetSelectorWithPrefixFromName("symbol"),
			"decimals":    encoding.GetSelectorWithPrefixFromName("decimals"),
			"totalSupply": encoding.GetSelectorWithPrefixFromName("totalSupply"),
		},
	}
}

// BalanceOf - Returns the number of tokens in owner's account.
func (erc20 ERC20) BalanceOf(ctx context.Context, account data.Felt, opts ...CallOption) (decimal.Decimal, error) {
	options := NewCallOptions(opts...)
	selector := erc20.selectors["balanceOf"]

	response, err := erc20.api.CallContract(ctx, options.block, erc20.contract.String(), selector, []string{
		account.Decimal().String(),
	})
	if err != nil {
		return decimal.Zero, err
	}
	if len(response.Result) < 2 {
		return decimal.Zero, errors.Errorf("invalid response for balanceOf: %v", response.Result)
	}

	balance := data.NewUint256(response.Result[0], response.Result[1])
	return balance.Decimal()
}

// Name - Returns the name of the token.
func (erc20 ERC20) Name(ctx context.Context, opts ...CallOption) (string, error) {
	options := NewCallOptions(opts...)
	selector := erc20.selectors["name"]

	response, err := erc20.api.CallContract(ctx, options.block, erc20.contract.String(), selector, []string{})
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for name: %v", response.Result)
	}

	return response.Result[0].ToAsciiString(), nil
}

// Symbol - Returns the ticker symbol of the token.
func (erc20 ERC20) Symbol(ctx context.Context, opts ...CallOption) (string, error) {
	options := NewCallOptions(opts...)
	selector := erc20.selectors["symbol"]

	response, err := erc20.api.CallContract(ctx, options.block, erc20.contract.String(), selector, []string{})
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for symbol: %v", response.Result)
	}

	return response.Result[0].ToAsciiString(), nil
}

// Decimals - Returns the number of decimals the token uses - e.g. 8 means to divide the token amount by 100000000 to get its user representation.
func (erc20 ERC20) Decimals(ctx context.Context, opts ...CallOption) (uint64, error) {
	options := NewCallOptions(opts...)
	selector := erc20.selectors["decimals"]

	response, err := erc20.api.CallContract(ctx, options.block, erc20.contract.String(), selector, []string{})
	if err != nil {
		return 0, err
	}
	if len(response.Result) < 1 {
		return 0, errors.Errorf("invalid response for decimals: %v", response.Result)
	}

	return response.Result[0].Uint64()
}

// TotalSupply - Returns the amount of tokens in existence.
func (erc20 ERC20) TotalSupply(ctx context.Context, opts ...CallOption) (decimal.Decimal, error) {
	options := NewCallOptions(opts...)
	selector := erc20.selectors["totalSupply"]

	response, err := erc20.api.CallContract(ctx, options.block, erc20.contract.String(), selector, []string{})
	if err != nil {
		return decimal.Zero, err
	}
	if len(response.Result) < 2 {
		return decimal.Zero, errors.Errorf("invalid response for totalSupply: %v", response.Result)
	}

	totalSupply := data.NewUint256(response.Result[0], response.Result[1])
	return totalSupply.Decimal()
}
