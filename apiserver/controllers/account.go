package controllers

import (
	"2corp/d2/apiserver/configs/constants/contextnames"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	accountID := c.Get(contextnames.AccountID)

	return c.String(http.StatusOK, fmt.Sprintf(`your account ID is "%v"`, accountID))
}
