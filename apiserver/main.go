package main

import (
	_ "2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.Company(e)

	e.Logger.Fatal(e.Start(":8080"))
}
