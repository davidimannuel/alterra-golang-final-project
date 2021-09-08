package user

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/user"
	"keep-remind-app/server/handlers"
	"keep-remind-app/server/handlers/user/request"
	"keep-remind-app/server/handlers/user/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	ContextUC *businesses.ContextUC
	Usecase   user.Usecase
}

func NewHandler(contextUC *businesses.ContextUC, uc user.Usecase) *Handler {
	return &Handler{
		ContextUC: contextUC,
		Usecase:   uc,
	}
}

func (h *Handler) Get(c echo.Context) error {

	return handlers.SendResponse(c, "ok", nil, nil, http.StatusOK)
}

func (h *Handler) Add(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), h.ContextUC.AppTimeout)
	defer cancel()
	req := new(request.JSON)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	data, err := h.Usecase.Add(ctx, req.ToDomain())
	res := response.JSON{
		Id:          data.Id,
		Username:    data.Username,
		CountryCode: data.CountryCode,
		Phone:       data.Phone,
		Email:       data.Email,
	}
	return handlers.SendResponse(c, res, nil, err, http.StatusOK)
}
