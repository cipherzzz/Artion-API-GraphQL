// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CanMintErc721 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) TokenBalance(contract common.Address, address common.Address) (*big.Int, error) {

	token, err := contracts.NewErc20Caller(contract, o.ftm)
	if err != nil {
		return nil, err
	}

	balance, error := token.BalanceOf(nil, address)
	if error != nil {
		return nil, error
	}

	return balance, nil
}
