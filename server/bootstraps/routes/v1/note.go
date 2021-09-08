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
	ContextUC   *businesses.ContextUC
}

func (route *Note) RegisterRoute() {
	r := route.RouterGroup.Group("/api/notes")
	repository := repo.NewPostgreSQLRepository(route.ContextUC.DB)
	usecase := uc.NewUsecase(repository)
	handler := handler.NewHandler(route.ContextUC, usecase)
	r.GET("", handler.Get)
	r.POST("", handler.Add)
}
