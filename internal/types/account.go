// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Account struct {
	Balance hexutil.Big `bson:"nonce"`
	Nonce   hexutil.Big `bson:"nonce"`
}
