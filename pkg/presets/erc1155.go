package presets

import (
	"context"
	"strconv"

	"github.com/dipdup-io/starknet-go-api/pkg/data"
	"github.com/dipdup-io/starknet-go-api/pkg/encoding"
	"github.com/dipdup-io/starknet-go-api/pkg/sequencer"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// ERC1155 -
type ERC1155 struct {
	api      sequencer.API
	contract data.Felt

	selectors map[string]string
}

// NewERC1155 -
func NewERC1155(api sequencer.API, contract data.Felt) ERC1155 {
	return ERC1155{
		api:      api,
		contract: contract,
		selectors: map[string]string{
			"balanceOf":        encoding.GetSelectorWithPrefixFromName("balanceOf"),
			"uri":              encoding.GetSelectorWithPrefixFromName("uri"),
			"isApprovedForAll": encoding.GetSelectorWithPrefixFromName("isApprovedForAll"),
			"balanceOfBatch":   encoding.GetSelectorWithPrefixFromName("balanceOfBatch"),
		},
	}
}

// BalanceOf - Returns the number of tokens in owner's account.
func (erc1155 ERC1155) BalanceOf(ctx context.Context, account data.Felt, tokenId data.Uint256, opts ...CallOption) (decimal.Decimal, error) {
	options := NewCallOptions(opts...)
	selector := erc1155.selectors["balanceOf"]

	calldata := []string{
		account.Decimal().String(),
	}
	calldata = append(calldata, tokenId.Calldata()...)

	response, err := erc1155.api.CallContract(ctx, options.block, erc1155.contract.String(), selector, calldata)
	if err != nil {
		return decimal.Zero, err
	}
	if len(response.Result) < 2 {
		return decimal.Zero, errors.Errorf("invalid response for balanceOf: %v", response.Result)
	}

	balance := data.NewUint256(response.Result[0], response.Result[1])
	return balance.Decimal()
}

// BalanceOfBatch - Get the balance of multiple account/token pairs.
func (erc1155 ERC1155) BalanceOfBatch(ctx context.Context, accounts []data.Felt, tokenIds []data.Uint256, opts ...CallOption) ([]decimal.Decimal, error) {
	options := NewCallOptions(opts...)
	selector := erc1155.selectors["balanceOfBatch"]

	accountsLen := strconv.FormatInt(int64(len(accounts)), 16)
	calldata := []string{
		encoding.AddHexPrefix(accountsLen),
	}
	for i := range accounts {
		calldata = append(calldata, accounts[i].Decimal().String())
	}
	tokenIdsLen := strconv.FormatInt(int64(len(tokenIds)), 16)
	calldata = append(calldata, encoding.AddHexPrefix(tokenIdsLen))
	for i := range tokenIds {
		calldata = append(calldata, tokenIds[i].Calldata()...)
	}

	response, err := erc1155.api.CallContract(ctx, options.block, erc1155.contract.String(), selector, calldata)
	if err != nil {
		return nil, err
	}
	if len(response.Result) < 1 {
		return nil, errors.Errorf("invalid response for balanceOfBatch: %v", response.Result)
	}

	balances := make([]decimal.Decimal, 0)
	for i := 1; i < len(response.Result); i += 2 {
		balance := data.NewUint256(response.Result[0], response.Result[1])
		value, err := balance.Decimal()
		if err != nil {
			return nil, err
		}
		balances = append(balances, value)
	}
	return balances, nil
}

// Uri - Returns the Uniform Resource Identifier (URI) for a token id
func (erc1155 ERC1155) Uri(ctx context.Context, tokenId data.Uint256, opts ...CallOption) (string, error) {
	options := NewCallOptions(opts...)
	selector := erc1155.selectors["uri"]

	response, err := erc1155.api.CallContract(ctx, options.block, erc1155.contract.String(), selector, tokenId.Calldata())
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for uri: %v", response.Result)
	}

	return response.Result[0].ToAsciiString(), nil
}

// IsApprovedForAll - Get whether operator is approved by account for all tokens.
func (erc1155 ERC1155) IsApprovedForAll(ctx context.Context, account, operator data.Felt, opts ...CallOption) (data.Felt, error) {
	options := NewCallOptions(opts...)
	selector := erc1155.selectors["isApprovedForAll"]

	calldata := []string{
		account.Decimal().String(),
		operator.Decimal().String(),
	}

	response, err := erc1155.api.CallContract(ctx, options.block, erc1155.contract.String(), selector, calldata)
	if err != nil {
		return "", err
	}
	if len(response.Result) < 1 {
		return "", errors.Errorf("invalid response for isApprovedForAll: %v", response.Result)
	}

	return response.Result[0], nil
}
