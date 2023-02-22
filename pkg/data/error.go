package data

import stdJSON "encoding/json"

// Error -
type Error struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    stdJSON.RawMessage `json:"data"`
}
