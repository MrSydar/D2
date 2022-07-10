package responses

import (
	"net/http"

	"github.com/labstack/echo"
)

func Created(c echo.Context, result interface{}) error {
	return c.JSON(
		http.StatusCreated,
		Body{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &echo.Map{"data": result},
		},
	)
}
