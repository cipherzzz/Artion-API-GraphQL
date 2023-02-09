package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// EstimateMintGas provides platform fee and gas estimation for new token minting
func (o *Opera) EstimateMintGas(user common.Address, contract common.Address, tokenUri string, royalty uint16) (platformFee *big.Int, gas uint64, err error) {
	contr, err := contracts.NewArtion(contract, o.ftm)
	if err != nil {
		return nil, 0, err
	}

	// load minting fee user the collection
	platformFee, err = contr.PlatformFee(&bind.CallOpts{
		From: user,
	})
	if err != nil {
		log.Debugf("contract %s does not define platformFee: %s", contract, err)
		platformFee = big.NewInt(0)
	}

	// construct minting transaction
	abi, err := contracts.ArtionMetaData.GetAbi()
	if err != nil {
		return nil, 0, err
	}

	//   function mint(address _to, string calldata _tokenUri) external payable
	// contract.mint(address _beneficiary, string _tokenUri, address _royaltyRecipient, uint256 _royaltyValue)
	data, err := abi.Pack("mint", user, tokenUri, user, big.NewInt(int64(royalty)))
	if err != nil {
		return nil, 0, err
	}

	fmt.Println("tokenUri", tokenUri)
	fmt.Println("user", user)
	fmt.Println("royalty", royalty)
	fmt.Println("data", data)

	// estimate minting gas
	// gas, err = o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	From:     user,
	// 	To:       &contract,
	// 	Value:    platformFee,
	// 	Gas:      0,
	// 	GasPrice: big.NewInt(0),
	// 	Data:     data,
	// })
	gas = 1000000
	return platformFee, gas, err
}
