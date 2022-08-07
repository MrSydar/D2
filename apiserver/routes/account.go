package routes

import (
	"2corp/d2/apiserver/configs/auth0"
	"2corp/d2/apiserver/controllers"
	d2middleware "2corp/d2/apiserver/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyAccount(e *echo.Echo) {
	e.GET("/account", controllers.GetAccount,
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(auth0.Secret)}),
		d2middleware.AssociateAccountWithRequest,
	)
}
