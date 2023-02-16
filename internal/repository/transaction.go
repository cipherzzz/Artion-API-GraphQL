// Package repository implements persistent data access and processing.
package repository

import "github.com/ethereum/go-ethereum/common"

// Estimate Gas
func (p *Proxy) TransactionStatus(hash common.Hash) (bool, error) {
	status, err := p.rpc.TransactionStatus(hash)
	if err != nil {
		return false, err
	}
	return status, nil
}
