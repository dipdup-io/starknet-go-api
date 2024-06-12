package data

// swagger:enum DAMode
/*
	ENUM(
		L1,
		L2
	)
*/
//go:generate go-enum --marshal --sql --values
type DAMode string
