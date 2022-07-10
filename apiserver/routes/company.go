package routes

import (
	"2corp/d2/apiserver/controllers"

	"github.com/labstack/echo"
)

func Company(e *echo.Echo) {
	e.POST("/company", controllers.CreateCompany)
	e.GET("/company/:id", controllers.GetCompany)
}
