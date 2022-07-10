package servererror

import (
	"2corp/d2/apiserver/responses"
	"net/http"

	"github.com/labstack/echo"
)

func Internal(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, responses.Default{Status: http.StatusInternalServerError, Message: "internal server error", Data: &echo.Map{"data": err.Error()}})
}
