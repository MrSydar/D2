package main

import (
	"2corp/d2/apiserver/configs"
	_ "2corp/d2/apiserver/configs"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	_, err := configs.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	e.Logger.Fatal(e.Start(":6000"))
}
