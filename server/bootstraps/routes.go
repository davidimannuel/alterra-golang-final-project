package bootstraps

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (boot *Bootstrap) RegisterRoute() {
	//test route
	boot.App.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Work!!")
	})
	// v1 routes
	apiV1 := boot.App.Group("/v1/api")
	// auth
	auth := apiV1.Group("/auth")
	auth.POST("/register", boot.AuthHandler.Register)
	auth.POST("/login", boot.AuthHandler.Login)

	user := apiV1.Group("/users")
	user.POST("", boot.UserHandler.Add)
}
