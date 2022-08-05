package controllers

import (
	"io"
	"net/http"

	"2corp/d2/apiserver/configs"

	"github.com/labstack/echo/v4"
)

func FetchJWTToken(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		msg := "code parameter was not provided"
		c.Logger().Error(msg)
		return c.String(http.StatusBadRequest, msg)
	}

	url := configs.Configs.Auth0.TokenFetchURL
	data := configs.Configs.Auth0.GetDataForTokenFetchWithCode(code)

	response, err := http.PostForm(url, data)
	if err != nil {
		msg := "failed to retrieve JWT token from Auth0 server"
		c.Logger().Error(msg + ": " + err.Error())
		return c.String(http.StatusServiceUnavailable, msg)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		msg := "failed to read response form Auth0 server"
		c.Logger().Error(msg + ": " + err.Error())
		return c.String(http.StatusInternalServerError, msg)
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusForbidden {
			msg := "got response from auth0: unauthorized"
			return c.String(http.StatusForbidden, msg)
		} else {
			msg := "got bad response from auth0"
			c.Logger().Error(msg + ": " + string(body))
			return c.String(http.StatusServiceUnavailable, msg)
		}
	}

	return c.String(http.StatusOK, string(body))
}
