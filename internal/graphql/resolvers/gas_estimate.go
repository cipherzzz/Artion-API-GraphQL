package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PayTokens provides list of tokens supported for payments on the marketplace
func (r *RootResolver) EstimateGas(args struct {
	To    common.Address
	From  common.Address
	Value hexutil.Big
	Data  string
}) (out types.GasEstimate, err error) {

	gas, err := repository.R().EstimateGas(args.To, args.From, *args.Value.ToInt(), args.Data)
	if err != nil {
		return types.GasEstimate{}, err
	}

	return types.GasEstimate{
		Gas: hexutil.Big(*big.NewInt(int64(gas))),
	}, nil
}
