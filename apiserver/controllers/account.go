package controllers

import (
	"2corp/d2/apiserver/configs/constants/contextnames"
	"2corp/d2/apiserver/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	accountID := c.Get(contextnames.AccountID)

	return responses.Message(c, http.StatusOK, fmt.Sprint(accountID))
}
