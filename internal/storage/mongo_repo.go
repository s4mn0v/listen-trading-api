package storage

import (
	"context"
	"os"
	"time"

	"github.com/s4mn0v/listen-trading-api/internal/models"
	"github.com/s4mn0v/listen-trading-api/logging/applogger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var TraderCol *mongo.Collection
var PositionsCol *mongo.Collection

func InitMongo() {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		applogger.Fatal("Error Mongo: %v", err)
	}

	MongoClient = client
	TraderCol = client.Database(dbName).Collection("traders")
	PositionsCol = client.Database(dbName).Collection("positions")
	applogger.Info("DB Mongo conectada: %s", dbName)
}

func SaveTraders(traders []models.TraderInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if len(traders) == 0 {
		return nil
	}

	var writes []mongo.WriteModel
	for _, t := range traders {
		filter := bson.M{"traderid": t.TraderId}
		model := mongo.NewReplaceOneModel().SetFilter(filter).SetReplacement(t).SetUpsert(true)
		writes = append(writes, model)
	}

	_, err := TraderCol.BulkWrite(ctx, writes)
	return err
}

func SavePositions(traderId string, positions []models.PositionInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, _ = PositionsCol.DeleteMany(ctx, bson.M{"traderId": traderId})

	if len(positions) == 0 {
		return nil
	}

	var list []interface{}
	for _, p := range positions {
		p.TraderId = traderId
		list = append(list, p)
	}
	_, err := PositionsCol.InsertMany(ctx, list)
	return err
}
