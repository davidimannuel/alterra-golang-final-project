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
	users := apiV1.Group("/users", middlewares.ContextManagement(boot.Configs.AppTimeout))
	boot.UserHandler.InitRoutes(users)
	// notes
	notes := apiV1.Group("/notes", middleware.JWTWithConfig(boot.Configs.JWT.Init()), middlewares.ContextManagement(boot.Configs.AppTimeout))
	notes.GET("", boot.NoteHandler.Get)
	notes.POST("", boot.NoteHandler.Add)
}
