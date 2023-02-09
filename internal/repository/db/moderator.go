package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// coModerators is the name of database collection.
	coModerators = "moderators"

	// fiModeratorAddress is the column storing the address of the user.
	fiModeratorAddress = "address"
)

func (sdb *SharedMongoDbBridge) IsModerator(address common.Address) (isMod bool, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coModerators)

	fmt.Println("moderators", address, sdb.dbName)

	filter := bson.D{{Key: fiModeratorAddress, Value: strings.ToLower(address.String())}}
	result := col.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
