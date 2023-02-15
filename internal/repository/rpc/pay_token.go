package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"artion-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
)

// ListPayTokens obtains list of tokens allowed for market payments from TokenRegistry contract
func (o *Opera) ListPayTokens() (payTokens []types.PayToken, err error) {
	// filterOps := bind.FilterOpts{
	// 	Context: context.Background(),
	// 	Start:   13817065,
	// 	End:     nil,
	// }
	// log.Debugf("ListPayTokens: %s", o.tokenRegistryContract)
	// itr, err := o.tokenRegistryContract.FilterTokenAdded(&filterOps)
	// if err != nil {
	// 	return nil, err
	// }

	addresses := []string{"0xf1277d1Ed8AD466beddF92ef448A132661956621", "0xEdE59D58d9B8061Ff7D22E629AB2afa01af496f4", "0x25C60d33451CD01350737F7809025DE5c2E10484"}

	for _, address := range addresses {
		payToken, err := o.getPayToken(common.HexToAddress(address))
		if err != nil {
			return nil, err
		}
		payTokens = append(payTokens, payToken)
	}
	return
}

func (o *Opera) getPayToken(address common.Address) (payToken types.PayToken, err error) {
	token, err := contracts.NewErc20(address, o.ftm)
	log.Debugf("getPayToken: %s", address.String(), token)
	if err != nil {
		return
	}
	payToken.Contract = address
	payToken.Name, err = token.Name(nil)
	if err != nil {
		return
	}
	payToken.Symbol, err = token.Symbol(nil)
	if err != nil {
		return
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		return
	}
	payToken.Decimals = int32(decimals)

	log.Debugf("address", &address)
	payToken.UnitPrice, err = o.GetPayTokenPrice(&address, nil)
	return
}
