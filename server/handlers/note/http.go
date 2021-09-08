package note

import (
	"context"
	"keep-remind-app/businesses"
	"keep-remind-app/businesses/note"
	"keep-remind-app/server/handlers"
	"keep-remind-app/server/handlers/note/request"
	"keep-remind-app/server/handlers/note/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Note struct {
	CtxUc   *businesses.ContextUC
	Usecase note.Usecase
}

func (h *Note) Get(c echo.Context) error {

	return handlers.SendResponse(c, "ok", nil, nil, http.StatusOK)
}

func (h *Note) Add(c echo.Context) error {
	req := new(request.Note)
	if err := c.Bind(req); err != nil {
		return handlers.SendBadResponse(c, err, http.StatusBadRequest)
	}
	data, err := h.Usecase.Add(context.Background(), req.ToDomain())
	res := response.Note{
		Id:         data.Id,
		Title:      data.Title,
		Note:       data.Title,
		ReminderAt: data.ReminderAt,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
	return handlers.SendResponse(c, res, nil, err, http.StatusOK)
}
