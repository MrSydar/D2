package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvNameMongoURI     = "MONGOURI"
	EnvNameDatabaseName = "DATABASENAME"
)

func EnvMongoURI() string {
	return os.Getenv(EnvNameMongoURI)
}

func EnvDatabaseName() string {
	return os.Getenv(EnvNameDatabaseName)
}

func verify() error {
	if EnvMongoURI() == "" {
		return fmt.Errorf("%v is empty", EnvNameMongoURI)
	}

	if EnvDatabaseName() == "" {
		return fmt.Errorf("%v is empty", EnvNameDatabaseName)
	}

	return nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	if err := verify(); err != nil {
		log.Fatalf("Failed to verify environment variables: %v", err)
	}
}
