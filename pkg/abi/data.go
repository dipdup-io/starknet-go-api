package abi

// abi types
const (
	FunctionType    = "function"
	L1HandlerType   = "l1_handler"
	ConstructorType = "constructor"
	EventType       = "event"
	StructType      = "struct"
)

// Abi -
type Abi struct {
	Functions   map[string]*FunctionItem `json:"-"`
	L1Handlers  map[string]*FunctionItem `json:"-"`
	Constructor map[string]*FunctionItem `json:"-"`
	Events      map[string]*EventItem    `json:"-"`
	Structs     map[string]*StructItem   `json:"-"`

	FunctionsBySelector   map[string]*FunctionItem `json:"-"`
	L1HandlersBySelector  map[string]*FunctionItem `json:"-"`
	ConstructorBySelector map[string]*FunctionItem `json:"-"`
	EventsBySelector      map[string]*EventItem    `json:"-"`
	StructsBySelector     map[string]*StructItem   `json:"-"`

	Names map[string]string `json:"-"`
}

// Type -
type Type struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// FunctionItem -
type FunctionItem struct {
	Type

	Inputs  []Type `json:"inputs"`
	Outputs []Type `json:"outputs"`
}

// EventItem -
type EventItem struct {
	Type

	Data []Type `json:"data"`
	Keys []Type `json:"keys"`
}

// StructItem -
type StructItem struct {
	Type

	Size    uint64   `json:"size"`
	Members []Member `json:"members"`
}

// Member -
type Member struct {
	Type

	Offset uint64 `json:"offset"`
}

// entrypoint types
const (
	EntrypointTypeExternal    = "EXTERNAL"
	EntrypointTypeConstructor = "CONSTRUCTOR"
	EntrypointTypeL1Handler   = "L1_HANDLER"
)
