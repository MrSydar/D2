package configs

import (
	"log"
	"reflect"
)

type IConfig interface {
	Init() error
	Verify() error
}

func init() {
	log.Println("Initializing configurations")

	reflectNames := reflect.ValueOf(Configs)

	for i := 0; i < reflectNames.NumField(); i++ {
		fieldValue := reflectNames.Field(i).Interface().(IConfig)
		fieldName := reflectNames.Type().Field(i).Name

		if fieldValue == nil {
			log.Fatalf("%v config is not set", fieldName)
		}

		if err := fieldValue.Init(); err != nil {
			log.Fatalf("failed to initialize %v config: %v", fieldName, err)
		}

		if err := fieldValue.Verify(); err != nil {
			log.Fatalf("failed to verify %v config: %v", fieldName, err)
		}
	}
}
