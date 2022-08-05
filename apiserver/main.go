package main

import (
	_ "2corp/d2/apiserver/configs"

	"2corp/d2/apiserver/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.ApplyCallback(e)

	e.Logger.Fatal(e.Start(":9000"))
}
