package data

const DefaultJSONRPC = "2.0"

// default block string names
const (
	Latest  = "latest"
	Pending = "pending"
)

// versions
const (
	Version0 = "0x0"
	Version1 = "0x1"
)

// abi types
const (
	AbiFunctionType    = "function"
	AbiL1HandlerType   = "l1_handler"
	AbiConstructorType = "constructor"
	AbiEventType       = "event"
	AbiStructType      = "struct"
)

// call types
const (
	CallTypeCall     = "CALL"
	CallTypeDelegate = "DELEGATE"
)

// entrypoint types
const (
	EntrypointTypeExternal    = "EXTERNAL"
	EntrypointTypeConstructor = "CONSTRUCTOR"
)

// statuses
const (
	StatusNotReceived  = "NOT_RECEIVED"
	StatusReceived     = "RECEIVED"
	StatusPending      = "PENDING"
	StatusRejected     = "REJECTED"
	StatusAcceptedOnL2 = "ACCEPTED_ON_L2"
	StatusAcceptedOnL1 = "ACCEPTED_ON_L1"
)
