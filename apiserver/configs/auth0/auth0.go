package auth0

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"2corp/d2/apiserver/configs/environment"
)

type Config struct {
	TokenFetchURL                string
	CallbackEndpoint             string
	Secret                       string
	GetDataForTokenFetchWithCode func(code string) url.Values
}

func (c *Config) Verify() error {
	return nil
}

func (c *Config) Init() error {
	log.Print("Initializing Auth0 variables")

	c.TokenFetchURL = fmt.Sprintf("https://%s/oauth/token", os.Getenv(environment.Configuration.VariableNames.Auth0Domain))
	c.CallbackEndpoint = os.Getenv(environment.Configuration.VariableNames.Auth0CallbackEndpoint)
	c.Secret = os.Getenv(environment.Configuration.VariableNames.Auth0ClientSecret)

	defaultData := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {os.Getenv(environment.Configuration.VariableNames.Auth0ClientID)},
		"client_secret": {os.Getenv(environment.Configuration.VariableNames.Auth0ClientSecret)},
		"redirect_uri":  {os.Getenv(environment.Configuration.VariableNames.Auth0CallbackURL)},
	}

	c.GetDataForTokenFetchWithCode = func(code string) url.Values {
		data := defaultData
		data["code"] = []string{code}
		return data
	}

	return nil
}

var Configuration Config = Config{}
