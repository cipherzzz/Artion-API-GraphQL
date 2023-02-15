// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type GasEstimate struct {
	Gas hexutil.Big `bson:"nonce"`
}
