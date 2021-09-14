package auth

import (
	"context"
	"keep-remind-app/businesses/auth"
	"keep-remind-app/configs"
	"keep-remind-app/server/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	configs *configs.Configs
	usecase auth.AuthUsecase
}

func NewAuthHandler(configs *configs.Configs, uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		configs: configs,
		usecase: uc,
	}
}

func (h *AuthHandler) InitRoutes(router *echo.Group) {
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
}

func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(RegisterUserRequest)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	token, err := h.usecase.Register(ctx, req.ToDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, token, nil)
}

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(Login)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	token, err := h.usecase.Login(ctx, req.ToDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, token, nil)
}
