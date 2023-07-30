package api

// Request -
type Request struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      uint64 `json:"id"`
}

// Response -
type Response[T any] struct {
	Version string `json:"jsonrpc"`
	Result  T      `json:"result,omitempty"`
	ID      uint64 `json:"id"`
	Error   *Error `json:"error,omitempty"`
}

// CallRequest -
type CallRequest struct {
	ContractAddress    string   `json:"contract_address"`
	EntrypointSelector string   `json:"entry_point_selector"`
	Calldata           []string `json:"calldata"`
}
