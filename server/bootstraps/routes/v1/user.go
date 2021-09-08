package v1

import (
	"keep-remind-app/businesses"
	uc "keep-remind-app/businesses/user"
	repo "keep-remind-app/drivers/repositories/user"
	handler "keep-remind-app/server/handlers/user"

	"github.com/labstack/echo/v4"
)

type User struct {
	RouterGroup *echo.Group
	ContextUC   *businesses.ContextUC
}

func (route *User) RegisterRoute() {
	r := route.RouterGroup.Group("/api/users")
	repository := repo.NewPostgreSQLRepository(route.ContextUC.DB)
	usecase := uc.NewUsecase(repository)
	handler := handler.NewHandler(route.ContextUC, usecase)
	r.GET("", handler.Get)
	r.POST("", handler.Add)
}
