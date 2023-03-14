package data

// Event -
type Event struct {
	Order       uint64   `json:"order"`
	FromAddress string   `json:"from_address"`
	Keys        []string `json:"keys"`
	Data        []string `json:"data"`
}

// Message -
type Message struct {
	Order       uint64   `json:"order"`
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Selector    Felt     `json:"selector"`
	Payload     []string `json:"payload"`
	Nonce       Felt     `json:"nonce"`
}
