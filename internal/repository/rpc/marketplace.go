package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// IMarketplaceContract defines single interface for all marketplace contract versions.
type IMarketplaceContract interface {
	GetPayTokenPrice(token *common.Address, block *big.Int) (*big.Int, error)
}

// GetPayTokenPrice extracts price of 1 whole pay token in USD in 6-decimals fixed point using Marketplace contract.
func (o *Opera) GetPayTokenPrice(token *common.Address, block *big.Int) (*big.Int, error) {
	//log.Debugf("GetPayTokenPrice: %s", o.payTokenPriceContract)
	log.Debugf("GetPayTokenPrice: %s", token)
	log.Debugf("GetPayTokenPrice: %s", block)

	return o.payTokenPriceContract.GetPayTokenPrice(token, block)
}
