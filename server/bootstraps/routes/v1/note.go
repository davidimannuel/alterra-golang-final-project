package v1

import (
	"alterra-golang-final-project/businesses"
	uc "alterra-golang-final-project/businesses/note"
	repo "alterra-golang-final-project/drivers/repositories/note"
	handler "alterra-golang-final-project/server/handlers/note"

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
