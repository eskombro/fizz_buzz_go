package fizzbuzz

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type FizzBuzzStats struct {
	Params FizzBuzzParams `json:"params" form:"params" param:"params"`
	Count  int            `json:"count" form:"count" param:"count"`
}

var db MongoDB

type MongoDB struct {
	client     *mongo.Client
	database   string
	collection string
	host       string
	port       string
}

func ConnectDB() error {
	db = MongoDB{
		database:   "fizzbuzzdb",
		collection: "stats",
		host:       "localhost",
		port:       "27017",
	}
	log.Println("Connecting to DB")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s", db.host, db.port),
	))
	db.client = cl
	if err != nil {
		return err
	}

	err = db.client.Ping(ctx, readpref.Primary())
	return err
}

func AddRequest(params FizzBuzzParams) {
	collection := db.client.Database(db.database).Collection(db.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stats := FizzBuzzStats{
		Params: params,
		Count:  1,
	}
	var result FizzBuzzStats
	r := collection.FindOne(ctx, bson.M{"params": stats.Params}).Decode(&result)
	if r == mongo.ErrNoDocuments {
		log.Println("Request doesn't exist. Adding it to DB")
		addRequestDB(ctx, collection, stats)
	} else {
		log.Println("Request exists, updating count:", result)
		updateRequestDB(ctx, collection, stats)
	}
}

func GetMostFrequentRequest() FizzBuzzStats {
	stats := FizzBuzzStats{}
	return stats
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
