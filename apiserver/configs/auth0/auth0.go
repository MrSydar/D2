package auth0

import (
	"2corp/d2/apiserver/configs/environment"
	"fmt"
	"log"
	"net/url"
	"os"
)

type Config struct {
	GenerateGetTokenURL func(code string) string
}

func (c *Config) Verify() error {
	// TODO: create separate tests

	return nil
}

func (c *Config) Init() error {
	log.Print("Building auth0 token url")

	rawAuth0GetTokenURL, err := url.Parse(os.Getenv(environment.Configuration.VariableNames.Auth0GetTokenURL))
	if err != nil {
		return fmt.Errorf("failed to parse auth0 get token url: %v", err)
	}

	rawAuth0GetTokenURLQuery := rawAuth0GetTokenURL.Query()
	rawAuth0GetTokenURLQuery.Set("grant_type", "authorization_code")
	rawAuth0GetTokenURLQuery.Set("client_id", os.Getenv(environment.Configuration.VariableNames.Auth0ClientID))
	rawAuth0GetTokenURLQuery.Set("client_secret", os.Getenv(environment.Configuration.VariableNames.Auth0ClientSecret))
	rawAuth0GetTokenURLQuery.Set("redirect_uri", os.Getenv(environment.Configuration.VariableNames.Auth0CallbackFullURI))

	rawAuth0GetTokenURL.RawQuery = rawAuth0GetTokenURLQuery.Encode()

	c.GenerateGetTokenURL = func(code string) string {
		auth0GetTokenURLQuery := rawAuth0GetTokenURLQuery
		auth0GetTokenURLQuery.Set("code", code)

		auth0GetTokenURL := *rawAuth0GetTokenURL
		auth0GetTokenURL.RawQuery = auth0GetTokenURLQuery.Encode()

		return fmt.Sprint(auth0GetTokenURL)
	}

	return nil
}

var Configuration Config = Config{}
