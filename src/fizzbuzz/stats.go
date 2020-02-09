package fizzbuzz

import (
	"context"
	"log"
	"time"

	"github.com/eskombro/fizz_buzz_go/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FizzBuzzStats struct {
	Params FizzBuzzParams `json:"params" form:"params" param:"params"`
	Count  int            `json:"count" form:"count" param:"count"`
}

func AddRequest(params FizzBuzzParams) {
	collection := db.Dbconf.Client.Database(db.Dbconf.Database).Collection(db.Dbconf.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stats := FizzBuzzStats{
		Params: params,
		Count:  1,
	}
	var result FizzBuzzStats
	r := collection.FindOne(ctx, bson.M{"params": stats.Params}).Decode(&result)
	if r == mongo.ErrNoDocuments {
		log.Println("Request doesn't exist. Adding to DB:", stats)
		addRequestDB(ctx, collection, stats)
	} else {
		result.Count++
		log.Println("Request exists, updating count:", result)
		updateRequestDB(ctx, collection, stats)
	}
}

func GetMostFrequentRequest() (FizzBuzzStats, error) {
	collection := db.Dbconf.Client.Database(db.Dbconf.Database).Collection(db.Dbconf.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stats := FizzBuzzStats{}
	err := collection.FindOne(
		ctx,
		bson.M{},
		&options.FindOneOptions{Sort: &bson.D{{"count", -1}}},
	).Decode(&stats)
	if err != nil {
		log.Println(err)
	}
	return stats, err
}

func addRequestDB(ctx context.Context, collection *mongo.Collection, stats FizzBuzzStats) {
	_, err := collection.InsertOne(ctx, stats)
	if err != nil {
		log.Println("Counldn't add document to DB")
		return
	}
	log.Println("Added document")
}

func updateRequestDB(ctx context.Context, collection *mongo.Collection, stats FizzBuzzStats) {
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"params": stats.Params},
		bson.D{{"$inc", bson.D{{"count", 1}}}},
	)
	if err != nil {
		log.Println(err)
		log.Println("Counldn't add document to DB")
		return
	}
	log.Println("Added document")
}
