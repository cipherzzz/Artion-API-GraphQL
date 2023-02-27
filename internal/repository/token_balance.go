// Package repository implements persistent data access and processing.
package repository

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Estimate Gas
func (p *Proxy) TokenBalance(contract common.Address, address common.Address) (*big.Int, error) {
	balance, err := p.rpc.TokenBalance(contract, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
