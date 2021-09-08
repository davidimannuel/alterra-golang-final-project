package bootstraps

import (
	v1 "keep-remind-app/server/bootstraps/routes/v1"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (boot *Bootstrap) RegisterRoute() {
	//test route
	boot.App.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Work!!")
	})

	// init handler

	apiV1 := boot.App.Group("/v1")
	noteRoutes := v1.Note{RouterGroup: apiV1, CtxUc: &boot.ContextUC}
	noteRoutes.RegisterRoute()
}
