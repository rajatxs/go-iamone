package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

// Opens new connection to database
func Open() (err error) {
	var opts *options.ClientOptions

	if client != nil {
		return nil
	}

	uri := os.Getenv("IAMONE_DB_CONN_URL")
	dbname := os.Getenv("IAMONE_DB_NAME")

	opts = options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.Background(), opts)
	db = client.Database(dbname)

	if err == nil {
		log.Println("database connected")
	}

	return err
}

// Closes the active connection of the database
func Close() error {
	if client == nil {
		return nil
	}

	return client.Disconnect(context.Background())
}
