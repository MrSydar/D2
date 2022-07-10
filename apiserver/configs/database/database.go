package database

import (
	env "2corp/d2/apiserver/configs/environment"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func connectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoURI := os.Getenv(env.Names.MongoUri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new client: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return client, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv(env.Names.Database)
	return client.Database(dbName).Collection(collectionName)
}

func Init() {
	log.Print("Initializing database connection")

	var err error
	DB, err = connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
