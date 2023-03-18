package presets

import (
	"context"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
	"github.com/dipdup-io/starknet-go-api/pkg/sequencer"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// ERC721 -
type ERC721 struct {
	api      sequencer.API
	contract data.Felt

	selectors map[string]string
}

// NewERC721 -
func NewERC721(api sequencer.API, contract data.Felt) ERC721 {
	return ERC721{
		api:      api,
		contract: contract,
		selectors: map[string]string{
			"balanceOf":        encoding.GetSelectorWithPrefixFromName("balanceOf"),
			"ownerOf":          encoding.GetSelectorWithPrefixFromName("ownerOf"),
			"getApproved":      encoding.GetSelectorWithPrefixFromName("getApproved"),
			"isApprovedForAll": encoding.GetSelectorWithPrefixFromName("isApprovedForAll"),
		},
	}
}

// BalanceOf - Returns the number of tokens in owner's account.
func (erc721 ERC721) BalanceOf(ctx context.Context, account data.Felt, opts ...CallOption) (decimal.Decimal, error) {
	options := NewCallOptions(opts...)
	selector := erc721.selectors["balanceOf"]

	response, err := erc721.api.CallContract(ctx, options.block, erc721.contract.String(), selector, []string{
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

// OwnerOf - Returns the owner of the tokenId token.
func (erc721 ERC721) OwnerOf(ctx context.Context, tokenId data.Uint256, opts ...CallOption) (data.Felt, error) {
	options := NewCallOptions(opts...)
	selector := erc721.selectors["ownerOf"]

	response, err := erc721.api.CallContract(ctx, options.block, erc721.contract.String(), selector, tokenId.Calldata())
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for ownerOf: %v", response.Result)
	}

	return response.Result[0], nil
}

// GetApproved - Returns the account approved for tokenId token.
func (erc721 ERC721) GetApproved(ctx context.Context, tokenId data.Uint256, opts ...CallOption) (data.Felt, error) {
	options := NewCallOptions(opts...)
	selector := erc721.selectors["getApproved"]

	response, err := erc721.api.CallContract(ctx, options.block, erc721.contract.String(), selector, tokenId.Calldata())
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for getApproved: %v", response.Result)
	}

	return response.Result[0], nil
}

// IsApprovedForAll - Returns if the operator is allowed to manage all of the assets of owner.
func (erc721 ERC721) IsApprovedForAll(ctx context.Context, owner, operator data.Felt, opts ...CallOption) (data.Felt, error) {
	options := NewCallOptions(opts...)
	selector := erc721.selectors["isApprovedForAll"]

	response, err := erc721.api.CallContract(ctx, options.block, erc721.contract.String(), selector, []string{
		owner.Decimal().String(),
		operator.Decimal().String(),
	})
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for isApprovedForAll: %v", response.Result)
	}

	return response.Result[0], nil
}
