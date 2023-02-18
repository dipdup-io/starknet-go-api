package sequencer

// Response -
type Response[T any] struct {
	Result T `json:"result"`
}
