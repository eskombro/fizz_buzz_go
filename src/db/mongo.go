package db

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
	Client     *mongo.Client
	Database   string
	Collection string
	host       string
	port       string
	username   string
	password   string
}

var Dbconf = MongoDB{
	Database:   "fizzbuzzdbTEST",
	Collection: "stats",
	host:       "localhost",
	port:       "27017",
	username:   "admin",
	password:   "secret",
}

func getEnvVariables() {
	envHost, exists := os.LookupEnv("MONGODB_HOST")
	if exists {
		Dbconf.host = envHost
	}
	envPort, exists := os.LookupEnv("MONGODB_PORT")
	if exists {
		Dbconf.port = envPort
	}
	envUsername, exists := os.LookupEnv("MONGODB_USERNAME")
	if exists {
		Dbconf.username = envUsername
	}
	envPassword, exists := os.LookupEnv("MONGODB_PASSWORD")
	if exists {
		Dbconf.password = envPassword
	}
}

func ConnectDB() error {
	getEnvVariables()
	log.Println("Connecting to DB")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s", Dbconf.host, Dbconf.port),
	).SetAuth(options.Credential{
		Username: Dbconf.username,
		Password: Dbconf.password,
	}))
	Dbconf.Client = cl
	if err != nil {
		log.Println(err)
		return err
	}
	err = Dbconf.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Connected to DB")
	return nil
}
