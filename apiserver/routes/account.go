package routes

import (
	"2corp/d2/apiserver/controllers"

	"github.com/labstack/echo/v4"
)

func ApplyAccount(e *echo.Echo) {
	e.GET("/account/:_id", controllers.GetAccount)
}
