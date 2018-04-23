package types

import crypto "github.com/tendermint/go-crypto"

// Standard Signature
type StdSignature struct {
	crypto.PubKey       `json:"pub_key"` // optional
	crypto.Signature    `json:"signature"`
	AccountSequence     int64 `json:"acc_sequence"`
	TransactionSequence int64 `json:"tx_sequence"`
}
