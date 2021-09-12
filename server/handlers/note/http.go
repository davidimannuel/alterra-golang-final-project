package note

import (
	"context"
	"keep-remind-app/businesses/note"
	"keep-remind-app/configs"
	"keep-remind-app/server/handlers"
	"keep-remind-app/server/handlers/note/request"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	configs *configs.Configs
	usecase note.Usecase
}

func NewHandler(configs *configs.Configs, uc note.Usecase) *Handler {
	return &Handler{
		configs: configs,
		usecase: uc,
	}
}

func (h *Handler) InitRoutes(router *echo.Group) {
	router.GET("", h.Get)
	router.POST("", h.Add)
}

func (h *Handler) Get(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	log.Println("Ctx", ctx)
	return handlers.SendSucessResponse(c, "ok", nil)
}

func (h *Handler) Add(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(request.JSON)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	_, err := h.usecase.Add(ctx, req.ToDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "success", nil)
}
