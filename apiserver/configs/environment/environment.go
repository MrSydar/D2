package environment

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func verify() error {
	reflectNames := reflect.ValueOf(Names)

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

func Init() {
	log.Print("Initializing environment variables")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	if err := verify(); err != nil {
		log.Fatalf("Failed to verify environment variables: %v", err)
	}
}
