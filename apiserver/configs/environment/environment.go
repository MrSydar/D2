package environment

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	VariableNames struct {
		MongoUri,
		Database,
		CompanyCollection,
		ItemCollection,
		PlaceCollection string
	}
}

func (c *Config) Verify() error {
	reflectNames := reflect.ValueOf(c.VariableNames)

	for i := 0; i < reflectNames.NumField(); i++ {
		fieldValue := reflectNames.Field(i).Interface().(string)

		if fieldValue == "" {
			fieldName := reflectNames.Type().Field(i).Name
			return fmt.Errorf("%v is not set", fieldName)
		}

		if os.Getenv(fieldValue) == "" {
			return fmt.Errorf("%v value is empty", reflectNames.Field(i))
		}
	}

	return nil
}

func (c *Config) Init() error {
	log.Print("Initializing environment variables")

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %v", err)
	}

	c.VariableNames.MongoUri = "MONGOURI"
	c.VariableNames.Database = "DATABASENAME"
	c.VariableNames.CompanyCollection = "COMPANYCOLLECTION"
	c.VariableNames.ItemCollection = "ITEMCOLLECTION"
	c.VariableNames.PlaceCollection = "PLACECOLLECTION"

	return nil
}

var Configuration Config = Config{}
