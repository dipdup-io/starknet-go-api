package data

import (
	stdJSON "encoding/json"
	"strconv"

	"github.com/dipdup-io/starknet-go-api/pkg/abi"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

// Class -
type Class struct {
	EntryPointsByType EntrypointsByType  `json:"entry_points_by_type"`
	ClassVersion      string             `json:"contract_class_version"`
	Abi               abi.Abi            `json:"-"`
	RawAbi            stdJSON.RawMessage `json:"abi"`
}

// GetAbi -
func (c *Class) GetAbi() (abi.Abi, error) {
	var abi abi.Abi
	switch c.ClassVersion {
	case "0.1.0":
		data := c.RawAbi
		for {
			s, err := strconv.Unquote(string(data))
			if err != nil {
				break
			}
			data = []byte(s)
		}

		c.RawAbi = []byte(data)
		if !json.Valid(c.RawAbi) {
			return abi, errors.Errorf("abi is not JSON: %s", data)
		}
	default:
	}

	err := json.Unmarshal(c.RawAbi, &abi)
	return abi, err
}

// Handler -
type Handler struct {
	Offset      uint64 `json:"-"`
	FunctionIdx uint64 `json:"function_idx,omitempty"`
	Selector    string `json:"selector"`
}

// UnmarshalJSON -
func (h *Handler) UnmarshalJSON(data []byte) error {
	type buf Handler
	if err := json.Unmarshal(data, (*buf)(h)); err != nil {
		return err
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if value, ok := m["offset"]; ok {
		switch t := value.(type) {
		case float64:
			h.Offset = uint64(t)
		case string:
			offset, err := strconv.ParseUint(t, 0, 64)
			if err != nil {
				return err
			}
			h.Offset = offset
		}
	}
	return nil
}

// EntrypointsByType -
type EntrypointsByType struct {
	CONSTRUCTOR []Handler `json:"CONSTRUCTOR"`
	EXTERNAL    []Handler `json:"EXTERNAL"`
	L1HANDLER   []Handler `json:"L1_HANDLER"`
}
