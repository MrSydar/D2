package configs

import "log"

func init() {
	log.Println("Initializing configurations")

	initEnvironmentVariables()
	initDatabase()
}
