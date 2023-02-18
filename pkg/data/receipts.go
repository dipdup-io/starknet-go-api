package data

// Event -
type Event struct {
	FromAddress string   `json:"from_address"`
	Keys        []string `json:"keys"`
	Data        []string `json:"data"`
}

// Message -
type Message struct {
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Selector    string   `json:"selector"`
	Payload     []string `json:"payload"`
	Nonce       string   `json:"nonce"`
}
