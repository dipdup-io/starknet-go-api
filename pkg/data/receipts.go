package data

// Event -
type Event struct {
	Order       uint64 `json:"order"`
	FromAddress string `json:"from_address"`
	Keys        []Felt `json:"keys"`
	Data        []Felt `json:"data"`
}

// Message -
type Message struct {
	Order       uint64 `json:"order"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Selector    Felt   `json:"selector"`
	Payload     []Felt `json:"payload"`
	Nonce       Felt   `json:"nonce"`
}
