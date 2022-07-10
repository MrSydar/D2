package configs

import (
	"2corp/d2/apiserver/configs/database"
	"2corp/d2/apiserver/configs/environment"
)

var Configs = struct {
	Env      *environment.Config
	Database *database.Config
}{
	Env:      &environment.Configuration,
	Database: &database.Configuration,
}
