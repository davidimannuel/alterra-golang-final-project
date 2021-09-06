package v1

import (
	"alterra-golang-final-project/config"
	"alterra-golang-final-project/server/handlers"

	"github.com/labstack/echo/v4"
)

type Note struct {
	RouterGroup *echo.Group
	Config      *config.Configs
}

func (route Note) RegisterRoute() {
	r := route.RouterGroup.Group("/api/notes")
	handler := handlers.Note{Configs: route.Config}
	r.GET("", handler.Get)
}
