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
	Version2 = "0x2"
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
	EntrypointTypeL1Handler   = "L1_HANDLER"
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

// length
const (
	AddressBytesLength = 32
)
