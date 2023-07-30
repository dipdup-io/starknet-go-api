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

// ERC721Metadata -
type ERC721Metadata struct {
	api      sequencer.API
	contract data.Felt

	selectors map[string]string
}

// NewERC721Metadata -
func NewERC721Metadata(api sequencer.API, contract data.Felt) ERC721Metadata {
	return ERC721Metadata{
		api:      api,
		contract: contract,
		selectors: map[string]string{
			"name":     encoding.GetSelectorWithPrefixFromName("name"),
			"symbol":   encoding.GetSelectorWithPrefixFromName("symbol"),
			"tokenURI": encoding.GetSelectorWithPrefixFromName("tokenURI"),
		},
	}
}

// Name - Returns the name of the token.
func (erc ERC721Metadata) Name(ctx context.Context, opts ...CallOption) (string, error) {
	options := NewCallOptions(opts...)
	selector := erc.selectors["name"]

	response, err := erc.api.CallContract(ctx, options.block, erc.contract.String(), selector, []string{})
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for name: %v", response.Result)
	}

	return response.Result[0].ToAsciiString(), nil
}

// Symbol - Returns the ticker symbol of the token.
func (erc ERC721Metadata) Symbol(ctx context.Context, opts ...CallOption) (string, error) {
	options := NewCallOptions(opts...)
	selector := erc.selectors["symbol"]

	response, err := erc.api.CallContract(ctx, options.block, erc.contract.String(), selector, []string{})
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for symbol: %v", response.Result)
	}

	return response.Result[0].ToAsciiString(), nil
}

// TokenUri - Returns the Uniform Resource Identifier (URI) for tokenID token. If the URI is not set for the tokenId, the return value will be empty string.
func (erc ERC721Metadata) TokenUri(ctx context.Context, tokenId data.Uint256, opts ...CallOption) (string, error) {
	options := NewCallOptions(opts...)
	selector := erc.selectors["tokenURI"]

	response, err := erc.api.CallContract(ctx, options.block, erc.contract.String(), selector, tokenId.Calldata())
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for token uri: %v", response.Result)
	}

	if response.Result[0].String() == "0x0" {
		return "", nil
	}

	return response.Result[0].ToAsciiString(), nil
}
