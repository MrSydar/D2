package success

import (
	"2corp/d2/apiserver/responses"
	"net/http"

	"github.com/labstack/echo"
)

func Created(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusCreated, responses.Default{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}
