package auth

import (
	"context"
	"keep-remind-app/businesses/auth"
	"keep-remind-app/configs"
	"keep-remind-app/server/handlers"
	"keep-remind-app/server/handlers/auth/request"
	"keep-remind-app/server/handlers/auth/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	configs *configs.Configs
	usecase auth.Usecase
}

func NewHandler(configs *configs.Configs, uc auth.Usecase) *Handler {
	return &Handler{
		configs: configs,
		usecase: uc,
	}
}

func (h *Handler) InitRoutes(router *echo.Group) {
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
}

func (h *Handler) Register(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(request.Register)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	data, err := h.usecase.Register(ctx, req.ToDomain())
	res := response.JWT{
		Token: data.JWTToken,
	}
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, res, nil)
}

func (h *Handler) Login(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(request.Login)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	data, err := h.usecase.Login(ctx, req.ToDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	res := response.JWT{
		Token: data.JWTToken,
	}
	return handlers.SendSucessResponse(c, res, nil)
}
