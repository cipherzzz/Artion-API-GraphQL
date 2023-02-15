package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PayTokens provides list of tokens supported for payments on the marketplace
func (r *RootResolver) Account(args struct {
	Address common.Address
}) (out types.Account, err error) {

	balance, err := repository.R().GetBalance(args.Address)
	if err != nil {
		return types.Account{}, err
	}

	return types.Account{
		Balance: hexutil.Big(*balance),
		Nonce:   hexutil.Big(*big.NewInt(5)),
	}, nil
}
