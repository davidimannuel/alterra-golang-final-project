package v1

import (
	"keep-remind-app/businesses"
	uc "keep-remind-app/businesses/note"
	repo "keep-remind-app/drivers/repositories/note"
	handler "keep-remind-app/server/handlers/note"

	"github.com/labstack/echo/v4"
)

type Note struct {
	RouterGroup *echo.Group
	CtxUc       *businesses.ContextUC
}

func (route Note) RegisterRoute() {
	r := route.RouterGroup.Group("/api/notes")
	repository := repo.NewRepository(route.CtxUc.DB)
	usecase := uc.NewUsecase(repository)
	handler := handler.Note{CtxUc: route.CtxUc, Usecase: usecase}
	r.GET("", handler.Get)
	r.POST("", handler.Add)
}
