package types

type Fee struct {
	// GasFeeCap is the maximum price that the message sender is willing to pay per unit of gas
	GasFeeCap uint64 `json:"gas_fee_cap" graphql:"gas_fee_cap"`
	// GasPremium is the price per unit of gas that the message sender is willing to pay
	GasPremium uint64 `json:"gas_premium" graphql:"gas_premium"`
	// GasLimit is the maximum amount of gas that a message’s execution should be allowed to consume on chain.
	GasLimit int64 `json:"gas_limit" graphql:"gas_limit"`
}
