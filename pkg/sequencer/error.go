package sequencer

import (
	stdJSON "encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Error -
type Error struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    stdJSON.RawMessage `json:"data"`
}

// Error -
func (e Error) Error() string {
	return fmt.Sprintf("code=%s message=%s data=%s", e.Code, e.Message, string(e.Data))
}

// errors
var (
	ErrRequest = errors.New("request error")
)
