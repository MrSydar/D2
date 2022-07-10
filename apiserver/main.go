package main

import (
	_ "2corp/d2/apiserver/configs"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
