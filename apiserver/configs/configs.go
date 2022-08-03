package configs

import (
	"2corp/d2/apiserver/configs/auth0"
	"2corp/d2/apiserver/configs/database"
	"2corp/d2/apiserver/configs/environment"
)

var Configs = struct {
	Env      *environment.Config
	Database *database.Config
	Auth0    *auth0.Config
}{
	Env:      &environment.Configuration,
	Database: &database.Configuration,
	Auth0:    &auth0.Configuration,
}
