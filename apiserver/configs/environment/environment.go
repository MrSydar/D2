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
		AccountCollection,
		ItemCollection,
		PlaceCollection,
		Auth0CallbackURL,
		Auth0CallbackEndpoint,
		Auth0Domain,
		Auth0ClientID,
		Auth0ClientSecret string
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

	c.VariableNames.MongoUri = "MONGO_URI"
	c.VariableNames.Database = "DATABASE_NAME"
	c.VariableNames.AccountCollection = "ACCOUNT_COLLECTION"
	c.VariableNames.ItemCollection = "ITEM_COLLECTION"
	c.VariableNames.PlaceCollection = "PLACE_COLLECTION"
	c.VariableNames.Auth0CallbackURL = "AUTH0_CALLBACK_URL"
	c.VariableNames.Auth0CallbackEndpoint = "AUTH0_CALLBACK_ENDPOINT"
	c.VariableNames.Auth0Domain = "AUTH0_DOMAIN"
	c.VariableNames.Auth0ClientID = "AUTH0_CLIENT_ID"
	c.VariableNames.Auth0ClientSecret = "AUTH0_CLIENT_SECRET"

	return nil
}

var Configuration Config = Config{}
