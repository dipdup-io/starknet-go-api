package api

import (
	stdJSON "encoding/json"
	"errors"
	"fmt"
)

// Error -
type Error struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    stdJSON.RawMessage `json:"data"`
}

// Error -
func (e Error) Error() string {
	return fmt.Sprintf("code=%d message=%s data=%s", e.Code, e.Message, string(e.Data))
}

// errors
var (
	ErrRequest = errors.New("request error")
)
