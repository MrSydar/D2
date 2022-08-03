package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func QueryParamValidationFailed(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		Body{
			Status:  http.StatusBadRequest,
			Message: "failed to validate query parameters",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}

func BodyValidationFailed(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		Body{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}

func FieldValidationFailed(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		Body{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body fields",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}
