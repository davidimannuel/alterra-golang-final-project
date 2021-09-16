package telegramuser

import (
	"context"
	"errors"
	telegramuser "keep-remind-app/businesses/telegramUser"
	"keep-remind-app/server/handlers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TelegramUserHandler struct {
	telegramUserUsecase telegramuser.TelegramUserUsecase
}

func NewTelegramUserHandler(telegramUserUsecase telegramuser.TelegramUserUsecase) *TelegramUserHandler {
	return &TelegramUserHandler{
		telegramUserUsecase: telegramUserUsecase,
	}
}

func (h *TelegramUserHandler) InitRoutes(router *echo.Group) {
	router.GET("/select", h.FindAll)
	router.POST("", h.Add)
	router.GET("/reqActivated/id/:id", h.ReqActivated)
}

func (h *TelegramUserHandler) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	res, err := h.telegramUserUsecase.FindAll(ctx, &telegramuser.TelegramUserParameter{})
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, fromDomains(res), nil)
}

func (h *TelegramUserHandler) Add(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(AddTelegramUserRequest)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	_, err := h.telegramUserUsecase.Add(ctx, req.toDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "ok", nil)
}

func (h *TelegramUserHandler) ReqActivated(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return handlers.SendBadResponse(c, errors.New("invalid param"), http.StatusBadRequest)
	}
	res, err := h.telegramUserUsecase.GenerateActivatedOTP(ctx, id)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, res, nil)
}
