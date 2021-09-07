package note

import (
	"alterra-golang-final-project/businesses"
	"alterra-golang-final-project/businesses/note"
	"alterra-golang-final-project/server/handlers"
	"alterra-golang-final-project/server/handlers/note/request"
	"alterra-golang-final-project/server/handlers/note/response"
	"context"
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
