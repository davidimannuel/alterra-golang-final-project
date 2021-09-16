package bootstraps

import (
	"keep-remind-app/server/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (boot *Bootstrap) RegisterRoute() {
	//test route
	boot.App.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Work!!")
	})
	// v1 routes
	apiV1 := boot.App.Group("/v1/api")
	// auth
	auth := apiV1.Group("/auth", middlewares.ContextManagement(boot.Configs.AppTimeout))
	boot.AuthHandler.InitRoutes(auth)
	// users
	users := apiV1.Group("/users", middleware.JWTWithConfig(boot.Configs.JWT.Init()), middlewares.ContextManagement(boot.Configs.AppTimeout))
	boot.UserHandler.InitRoutes(users)
	// telegramUsers
	telegramUsers := apiV1.Group("/telegramUsers", middleware.JWTWithConfig(boot.Configs.JWT.Init()), middlewares.ContextManagement(boot.Configs.AppTimeout))
	boot.TelegramUserhandler.InitRoutes(telegramUsers)
	// labels
	labels := apiV1.Group("/labels", middleware.JWTWithConfig(boot.Configs.JWT.Init()), middlewares.ContextManagement(boot.Configs.AppTimeout))
	boot.LabelHandler.InitRoutes(labels)
	// notes
	notes := apiV1.Group("/notes", middleware.JWTWithConfig(boot.Configs.JWT.Init()), middlewares.ContextManagement(boot.Configs.AppTimeout))
	boot.NoteHandler.InitRoutes(notes)
}
