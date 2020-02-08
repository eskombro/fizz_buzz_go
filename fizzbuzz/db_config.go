package fizzbuzz

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client     *mongo.Client
	database   string
	collection string
	host       string
	port       string
	username   string
	password   string
}

var db = MongoDB{
	database:   "fizzbuzzdbTEST",
	collection: "stats",
	host:       "localhost",
	port:       "27017",
	username:   "admin",
	password:   "secret",
}

func getEnvVariables() {
	envHost, exists := os.LookupEnv("MONGODB_HOST")
	if exists {
		db.host = envHost
	}
	envPort, exists := os.LookupEnv("MONGODB_PORT")
	if exists {
		db.port = envPort
	}
	envUsername, exists := os.LookupEnv("MONGODB_USERNAME")
	if exists {
		db.username = envUsername
	}
	envPassword, exists := os.LookupEnv("MONGODB_PASSWORD")
	if exists {
		db.password = envPassword
	}
}

func ConnectDB() error {
	getEnvVariables()
	log.Println("Connecting to DB")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s", db.host, db.port),
	).SetAuth(options.Credential{
		Username: db.username,
		Password: db.password,
	}))
	db.client = cl
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Connected to DB")
	return nil
}
