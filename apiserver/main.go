package main

import (
	_ "2corp/d2/apiserver/configs"

	"2corp/d2/apiserver/routes"

	"github.com/labstack/echo/v4"
)

// AUTH URL:
// https://dev-tslb5vli.us.auth0.com/authorize?response_type=code&client_id=FXMNvdmbbyoy7EyvmwaMGRUdyuUFePwk&redirect_uri=http://localhost:9000/callback&scope=openid%20email&state=STATE

func main() {
	e := echo.New()

	routes.ApplyCallback(e)
	routes.ApplyAccount(e)

	e.Logger.Fatal(e.Start(":9000"))
}
