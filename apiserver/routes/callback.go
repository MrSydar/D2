package routes

import (
	"2corp/d2/apiserver/configs/auth0"
	"2corp/d2/apiserver/controllers"

	"github.com/labstack/echo/v4"
)

func ApplyCallback(e *echo.Echo) {
	e.GET(auth0.CallbackEndpoint, controllers.FetchJWTToken)
}
