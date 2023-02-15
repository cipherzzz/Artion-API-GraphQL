package resolvers

import (
	"artion-api-graphql/internal/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PayTokens provides list of tokens supported for payments on the marketplace
func (r *RootResolver) GasPrice() (out types.GasPrice, err error) {
	return types.GasPrice{
		Price: hexutil.Big(*big.NewInt(5)),
	}, nil
}
