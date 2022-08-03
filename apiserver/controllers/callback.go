package controllers

import (
	"2corp/d2/apiserver/configs"
	"2corp/d2/apiserver/responses"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetJWTToken(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		err := errors.New("code parameter was not provided")
		c.Logger().Error(err)
		return responses.QueryParamValidationFailed(c, err)
	}

	response, err := http.Post(configs.Configs.Auth0.GenerateGetTokenURL(code), "application/x-www-form-urlencoded", nil)
	if err != nil {
		err = fmt.Errorf("failed to ", err)
		return responses.InternalServerError(c)
	}

	return nil
}
