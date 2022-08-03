package main

import (
	_ "2corp/d2/apiserver/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// config := middleware.JWTConfig{
	// 	KeyFunc: ,
	// }

	// e.Use(middleware.JWTWithConfig(config))

	// routes.ApplyAccount(e)

	e.Logger.Fatal(e.Start(":9000"))
}
