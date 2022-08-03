package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(
		http.StatusInternalServerError,
		Body{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    &echo.Map{"data": err.Error()},
		})
}

func NotFound(c echo.Context, err error) error {
	return c.JSON(
		http.StatusNotFound,
		Body{
			Status:  http.StatusNotFound,
			Message: "requested resource was not found",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}
