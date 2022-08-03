package routes

import (
	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/controllers"
	"os"

	"github.com/labstack/echo/v4"
)

func ApplyCallback(e *echo.Echo) {
	callbackURI := os.Getenv(configs.Configs.Env.VariableNames.Auth0CallbackFullURI)
	e.GET("/"+callbackURI, controllers.GetJWTToken)
}
