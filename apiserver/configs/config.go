package configs

import (
	"2corp/d2/apiserver/configs/database"
	"2corp/d2/apiserver/configs/environment"
	"log"
)

func init() {
	log.Println("Initializing configurations")

	environment.Init()
	database.Init()
}
