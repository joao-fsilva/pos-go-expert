package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection   *mongo.Collection
	openAuctions map[string]auction_entity.AuctionStatus
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	err = ar.autoClose(ctx)
	if err != nil {
		return nil
	}

	return nil
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}

func (ar *AuctionRepository) autoClose(ctx context.Context) error {
	openAuctions, err := ar.FindOpenAuctions(ctx)
	if err != nil {
		return err
	}

	for _, auctionEntity := range openAuctions {
		auctionEndTime := auctionEntity.Timestamp.Add(getAuctionInterval())
		timeUntilClose := time.Until(auctionEndTime)

		if timeUntilClose <= 0 {
			err := ar.closeAuction(ctx, auctionEntity)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed to close auction %s immediately", auctionEntity.Id), err)
			}
			continue
		}

		_, okStatus := ar.openAuctions[auctionEntity.Id]
		if okStatus {
			continue
		}

		timer := time.NewTimer(timeUntilClose)

		go func(auction auction_entity.Auction) {
			<-timer.C
			err := ar.closeAuction(ctx, auction)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed to close auction %s automatically", auction.Id), err)
			}
		}(auctionEntity)
	}

	return nil
}

func (ar *AuctionRepository) closeAuction(ctx context.Context, auctionEntity auction_entity.Auction) error {
	filter := bson.M{"_id": auctionEntity.Id}
	update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

	_, err := ar.Collection.UpdateOne(
		ctx,
		filter,
		update,
		options.Update().SetUpsert(false),
	)

	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Auction %s closed automatically", auctionEntity.Id))

	return nil
}