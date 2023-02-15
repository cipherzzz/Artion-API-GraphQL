// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// CanMintErc721 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) TokenBalance(contract common.Address, address common.Address) (uint64, error) {

	token, err := contracts.NewErc20Caller(contract, o.ftm)
	if err != nil {
		return 0, err
	}
	balance, error := token.BalanceOf(&bind.CallOpts{}, address)
	if error != nil {
		return 0, error
	}
	return uint64(balance.Uint64()), nil
}
