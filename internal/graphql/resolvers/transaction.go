package resolvers

import (
	"artion-api-graphql/internal/repository"

	"github.com/ethereum/go-ethereum/common"
)

// PayTokens provides list of tokens supported for payments on the marketplace
func (r *RootResolver) TransactionPending(args struct {
	Hash string
}) (out bool, err error) {

	var hsh common.Hash
	err = hsh.UnmarshalGraphQL(args.Hash)
	if err != nil {
		return false, err
	}

	status, err := repository.R().TransactionStatus(hsh)
	if err != nil {
		return false, err
	}

	return status, nil
}
