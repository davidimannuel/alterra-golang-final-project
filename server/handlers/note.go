package handlers

import (
	"alterra-golang-final-project/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Note struct {
	Configs *config.Configs
}

func (h *Note) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
