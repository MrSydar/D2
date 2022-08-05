package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	env "2corp/d2/apiserver/configs/environment"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoURI := os.Getenv(env.Configuration.VariableNames.MongoUri)
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

func (c *Config) getCollection(collectionName string) *mongo.Collection {
	return c.DB.Database(c.dbName).Collection(collectionName)
}

type Config struct {
	DB     *mongo.Client
	dbName string

	Collections struct {
		Items,
		Places,
		Accounts *mongo.Collection
	}
}

func (c *Config) Verify() error {
	reflectNames := reflect.ValueOf(c.Collections)

	for i := 0; i < reflectNames.NumField(); i++ {
		fieldValue := reflectNames.Field(i).Interface().(*mongo.Collection)

		if fieldValue == nil {
			fieldName := reflectNames.Type().Field(i).Name
			return fmt.Errorf("%v is not set", fieldName)
		}
	}

	return nil
}

func (c *Config) Init() error {
	log.Print("Initializing database connection")

	var err error
	c.DB, err = connectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	c.dbName = os.Getenv(env.Configuration.VariableNames.Database)

	c.Collections.Accounts = c.getCollection(env.Configuration.VariableNames.AccountCollection)
	c.Collections.Items = c.getCollection(env.Configuration.VariableNames.ItemCollection)
	c.Collections.Places = c.getCollection(env.Configuration.VariableNames.PlaceCollection)

	return nil
}

var Configuration Config = Config{}
