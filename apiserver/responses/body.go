package responses

import (
	"github.com/labstack/echo"
)

type Body struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}
