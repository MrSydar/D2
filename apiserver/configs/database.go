package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(EnvDatabaseName()).Collection(collectionName)
}

var DB *mongo.Client

func connectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new client: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return client, nil
}

func initDatabase() {
	log.Print("Initializing database connection")

	var err error
	DB, err = connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
