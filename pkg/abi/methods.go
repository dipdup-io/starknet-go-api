package abi

import "github.com/dipdup-io/starknet-go-api/pkg/encoding"

// GetFunctionBySelector - receives function's ABI by selector. Selector may has prefix 0x and left-padding zeroes.
func (a *Abi) GetFunctionBySelector(selector string) (*FunctionItem, bool) {
	trimmed := encoding.TrimHex(selector)
	val, ok := a.FunctionsBySelector[trimmed]
	return val, ok
}

// GetL1HandlerBySelector - receives l1 handler's ABI by selector. Selector may has prefix 0x and left-padding zeroes.
func (a *Abi) GetL1HandlerBySelector(selector string) (*FunctionItem, bool) {
	trimmed := encoding.TrimHex(selector)
	val, ok := a.L1HandlersBySelector[trimmed]
	return val, ok
}

// GetConstructorBySelector - receives constructor's ABI by selector. Selector may has prefix 0x and left-padding zeroes.
func (a *Abi) GetConstructorBySelector(selector string) (*FunctionItem, bool) {
	trimmed := encoding.TrimHex(selector)
	val, ok := a.ConstructorBySelector[trimmed]
	return val, ok
}

// GetEventBySelector - receives event's ABI by selector. Selector may has prefix 0x and left-padding zeroes.
func (a *Abi) GetEventBySelector(selector string) (*EventItem, bool) {
	trimmed := encoding.TrimHex(selector)
	val, ok := a.EventsBySelector[trimmed]
	return val, ok
}

// GetStructBySelector - receives struct's ABI by selector. Selector may has prefix 0x and left-padding zeroes.
func (a *Abi) GetStructBySelector(selector string) (*StructItem, bool) {
	trimmed := encoding.TrimHex(selector)
	val, ok := a.StructsBySelector[trimmed]
	return val, ok
}
