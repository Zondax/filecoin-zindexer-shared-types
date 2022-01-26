package types

import "time"

type MempoolTransaction struct {
	FirstTimeSeen time.Time `json:"first_seen" graphql:"first_time_seen"`
	LastTimeSeen  time.Time `json:"last_seen"  graphql:"last_time_seen"`

	// Nonce is a nonce to prevent replay.
	Nonce uint64 `json:"nonce" graphql:"nonce"`
	// Value contains the value amount
	Value string `json:"value" graphql:"value"`
	// Fee is an optional fee that the sender commits to pay to execute this transaction.
	Fee `gorm:"embedded"`
	// MethodName is the string version of the method that should be called.
	MethodName string `json:"method_name" graphql:"method_name"`
	// Method num
	Method uint64 `gorm:"-" graphql:"# ignore\n"`
	// TxHash
	TxHash string `json:"tx_hash" gorm:"uniqueIndex" graphql:"tx_hash"`
	// From
	TxFrom string `json:"tx_from" graphql:"tx_from"`
	// To
	TxTo string `json:"tx_to" graphql:"tx_to"`
}
