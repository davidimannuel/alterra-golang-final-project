package user

import (
	"context"
	"keep-remind-app/businesses/user"
	"keep-remind-app/configs"
	"keep-remind-app/server/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	configs *configs.Configs
	usecase user.UserUsecase
}

func NewUserHandler(configs *configs.Configs, uc user.UserUsecase) *UserHandler {
	return &UserHandler{
		configs: configs,
		usecase: uc,
	}
}

func (h *UserHandler) InitRoutes(router *echo.Group) {
	router.GET("/profile", h.Profile)
}

func (h *UserHandler) Profile(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	param := new(user.UserParameter)
	param.ID = ctx.Value("user_id").(int)
	res, err := h.usecase.FindByID(ctx, param)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusNotFound)
	}
	return handlers.SendSucessResponse(c, FromDomain(&res), nil)
}
