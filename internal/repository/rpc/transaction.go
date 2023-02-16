// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// MustTransactionData provides call data of the given transaction, if available, or empty slice in other cases.
func (o *Opera) TransactionStatus(tx common.Hash) (bool, error) {
	_, pending, err := o.ftm.TransactionByHash(context.Background(), tx)
	if err != nil {
		log.Errorf("transaction %s detail unknown; %s", tx.String(), err.Error())
		return false, err
	}

	if pending {
		return true, nil
	}

	return false, nil
}
