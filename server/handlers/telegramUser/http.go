package telegramuser

import (
	"context"
	telegramuser "keep-remind-app/businesses/telegramUser"
	"keep-remind-app/server/handlers"
	"net/http"

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
