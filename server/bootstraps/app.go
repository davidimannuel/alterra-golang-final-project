package bootstraps

import (
	"keep-remind-app/businesses"

	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	App       *echo.Echo
	ContextUC businesses.ContextUC
}
