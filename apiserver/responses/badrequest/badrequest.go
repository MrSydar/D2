package badrequest

import (
	"2corp/d2/apiserver/responses"
	"net/http"

	"github.com/labstack/echo"
)

func BodyValidationFailed(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		responses.Default{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}

func FieldValidationFailed(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		responses.Default{
			Status:  http.StatusBadRequest,
			Message: "failed to validate request body fields",
			Data:    &echo.Map{"data": err.Error()},
		},
	)
}
