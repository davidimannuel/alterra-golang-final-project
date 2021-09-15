package user

import (
	"context"
	"keep-remind-app/businesses/user"
	"keep-remind-app/server/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) InitRoutes(router *echo.Group) {
	router.GET("/profile", h.Profile)
	router.PUT("/edit", h.Edit)
}

func (h *UserHandler) Profile(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	param := new(user.UserParameter)
	param.ID = ctx.Value("user_id").(int)
	res, err := h.userUsecase.FindOne(ctx, param)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusNotFound)
	}
	return handlers.SendSucessResponse(c, FromDomain(&res), nil)
}

func (h *UserHandler) Edit(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(EditUserRequest)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	userEdit := req.ToDomain()
	userEdit.ID = ctx.Value("user_id").(int)
	err := h.userUsecase.Edit(ctx, userEdit)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusNotFound)
	}

	return handlers.SendSucessResponse(c, "ok", nil)
}
