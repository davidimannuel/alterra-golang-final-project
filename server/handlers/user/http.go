package user

import (
	"context"
	"keep-remind-app/businesses/user"
	"keep-remind-app/configs"
	"keep-remind-app/server/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	configs *configs.Configs
	usecase user.Usecase
}

func NewHandler(configs *configs.Configs, uc user.Usecase) *Handler {
	return &Handler{
		configs: configs,
		usecase: uc,
	}
}

func (h *Handler) InitRoutes(router *echo.Group) {
	router.GET("/profile", h.Profile)
}

func (h *Handler) Profile(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	res, err := h.usecase.FindByID(ctx, user.Parameter{
		ID: ctx.Value("user_id").(int),
	})
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusNotFound)
	}
	return handlers.SendSucessResponse(c, FromDomain(&res), nil)
}
