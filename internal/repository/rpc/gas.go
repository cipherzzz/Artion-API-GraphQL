// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// CanMintErc721 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) EstimateGas(to common.Address, from common.Address, value big.Int, data []byte) (uint64, error) {

	// try to estimate the call
	gas, err := o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  from,
		To:    &to,
		Data:  data,
		Value: &value,
	})
	if err != nil {
		return 0, nil
	}

	return gas, nil
}
