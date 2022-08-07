package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	auth0Config "2corp/d2/apiserver/configs/auth0"
	"2corp/d2/apiserver/configs/log"

	"github.com/labstack/echo/v4"
)

func FetchJWTToken(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		msg := "code parameter was not provided"
		log.Logger.Error(msg)
		return c.String(http.StatusBadRequest, msg)
	}

	url := auth0Config.TokenFetchURL
	data := auth0Config.GetDataForTokenFetchWithCode(code)

	response, err := http.PostForm(url, data)
	if err != nil {
		msg := "failed to retrieve JWT token from Auth0 server"
		log.Logger.Error(msg + ": " + err.Error())
		return c.String(http.StatusServiceUnavailable, msg)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		msg := "failed to read response form Auth0 server"
		log.Logger.Error(msg + ": " + err.Error())
		return c.String(http.StatusInternalServerError, msg)
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusForbidden {
			msg := "got response from auth0: unauthorized"
			return c.String(http.StatusForbidden, msg)
		} else {
			msg := "got bad response from auth0"
			log.Logger.Error(msg + ": " + string(body))
			return c.String(http.StatusServiceUnavailable, msg)
		}
	}

	fieldsToCheck := struct {
		Scope string `json:"scope"`
	}{}

	if err := json.Unmarshal(body, &fieldsToCheck); err != nil {
		msg := "failed to unmarshal body for field check"
		log.Logger.Errorf(msg + ": " + string(body))
		return c.String(http.StatusInternalServerError, msg)
	}

	if !strings.Contains(fieldsToCheck.Scope, "email") {
		msg := `"email" scope is required`
		return c.String(http.StatusBadRequest, msg)
	}

	// TODO: return errors in `{"error": "info"}` format

	return c.String(http.StatusOK, string(body))
}
