package label

import (
	"context"
	"errors"
	"keep-remind-app/businesses/label"
	"keep-remind-app/server/handlers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LabelHandler struct {
	labelUsecase label.LabelUsecase
}

func NewLabelHandler(labelUsecase label.LabelUsecase) *LabelHandler {
	return &LabelHandler{
		labelUsecase: labelUsecase,
	}
}

func (h *LabelHandler) InitRoutes(router *echo.Group) {
	router.GET("/select", h.FindAll)
	router.GET("", h.FindAllPagination)
	router.POST("", h.Add)
	router.PUT("/id/:id", h.Edit)
	router.DELETE("/id/:id", h.Delete)
}

func (h *LabelHandler) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	param := new(label.LabelParameter)
	res, err := h.labelUsecase.FindAll(ctx, param)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusInternalServerError)
	}
	return handlers.SendSucessResponse(c, FromDomains(res), nil)
}

func (h *LabelHandler) FindAllPagination(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	param := new(label.LabelParameter)
	param.PerPage, _ = strconv.Atoi(c.QueryParam("per_page"))
	param.Page, _ = strconv.Atoi(c.QueryParam("page"))
	res, meta, err := h.labelUsecase.FindAllPagination(ctx, param)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusInternalServerError)
	}
	return handlers.SendSucessResponse(c, FromDomains(res), handlers.PageInfo(meta))
}

func (h *LabelHandler) Add(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	req := new(AddLabelRequest)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	_, err := h.labelUsecase.Add(ctx, req.toDomain())
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "ok", nil)
}

func (h *LabelHandler) Edit(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return handlers.SendBadResponse(c, errors.New("invalid param"), http.StatusBadRequest)
	}
	req := new(EditLabelRequest)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	labelEdit := req.toDomain()
	labelEdit.ID = id
	err := h.labelUsecase.Edit(ctx, labelEdit)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "ok", nil)
}

func (h *LabelHandler) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return handlers.SendBadResponse(c, errors.New("invalid param"), http.StatusBadRequest)
	}
	err := h.labelUsecase.Delete(ctx, id)
	if err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	return handlers.SendSucessResponse(c, "ok", nil)
}
