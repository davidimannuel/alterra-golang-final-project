package bootstraps

import (
	"alterra-golang-final-project/businesses"

	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	App       *echo.Echo
	ContextUC businesses.ContextUC
}
