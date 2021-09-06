package routes

import (
	"net/http"

	"alterra-golang-final-project/config"
	v1 "alterra-golang-final-project/server/routes/v1"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo    *echo.Echo
	Configs *config.Configs
}

func (r *Router) RegisterRoute() {
	//test route
	r.Echo.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Work!!")
	})

	apiV1 := r.Echo.Group("/v1")
	noteRoutes := v1.Note{RouterGroup: apiV1}
	noteRoutes.RegisterRoute()
}
