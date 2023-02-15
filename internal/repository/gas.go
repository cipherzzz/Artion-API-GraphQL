// Package repository implements persistent data access and processing.
package repository

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Estimate Gas
func (p *Proxy) EstimateGas(to common.Address, from common.Address, value big.Int,
	data string) (uint64, error) {
	gas, err := p.rpc.EstimateGas(to, from, value, []byte(data))
	if err != nil {
		return 0, err
	}
	return gas, nil
}
