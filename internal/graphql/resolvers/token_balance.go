package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PayTokens provides list of tokens supported for payments on the marketplace
func (r *RootResolver) ErcTokenBalance(args struct {
	Contract common.Address
	Account  common.Address
}) (out types.TokenBalance, err error) {

	balance, err := repository.R().TokenBalance(args.Contract, args.Account)
	if err != nil {
		return types.TokenBalance{}, err
	}

	return types.TokenBalance{
		Balance: hexutil.Big(*big.NewInt(int64(balance))),
	}, nil
}
