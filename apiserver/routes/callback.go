package routes

import (
	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/controllers"

	"github.com/labstack/echo/v4"
)

func ApplyCallback(e *echo.Echo) {
	e.GET(configs.Configs.Auth0.CallbackEndpoint, controllers.FetchJWTToken)
}
