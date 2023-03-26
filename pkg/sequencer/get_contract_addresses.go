package sequencer

import (
	"context"
)

// ContractAddresses -
type ContractAddresses struct {
	Starknet             string `json:"Starknet"`
	GpsStatementVerifier string `json:"GpsStatementVerifier"`
}

// GetContractAddresses -
func (api API) GetContractAddresses(ctx context.Context) (response ContractAddresses, err error) {
	err = api.getFromFeederGateway(ctx, "get_contract_addresses", "", nil, &response)
	return
}
