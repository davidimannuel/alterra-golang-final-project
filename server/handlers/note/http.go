package note

import (
	"bytes"
	"context"
	"io"
	"keep-remind-app/businesses/note"
	"keep-remind-app/configs"
	"keep-remind-app/server/handlers"
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
	router.POST("/image", h.AddWithImage)
}

func (h *Handler) Get(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	res, err := h.usecase.FindAll(ctx, note.Parameter{})
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusInternalServerError)
	}
	return handlers.SendSucessResponse(c, FromDomains(res), nil)
}

func (h *Handler) Add(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(AddNote)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	_, err := h.usecase.Add(ctx, req.ToDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "success", nil)
}

func (h *Handler) AddWithImage(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	title := c.Request().FormValue("title")
	file, _, err := c.Request().FormFile("note_image")
	defer file.Close()
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	_, err = h.usecase.AddWithImageBytes(ctx, title, buf.Bytes())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "ok", nil)
}
