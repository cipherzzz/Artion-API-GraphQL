// Package repository implements persistent data access and processing.
package repository

import (
	"github.com/ethereum/go-ethereum/common"
)

// Estimate Gas
func (p *Proxy) TokenBalance(contract common.Address, address common.Address) (uint64, error) {
	balance, err := p.rpc.TokenBalance(contract, address)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
