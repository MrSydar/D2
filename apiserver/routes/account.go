package routes

import (
	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyAccount(e *echo.Echo) {
	e.GET("/account", controllers.GetAccount, middleware.JWT([]byte(configs.Configs.Auth0.Secret)))
}
