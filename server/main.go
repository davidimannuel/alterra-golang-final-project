package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Work!!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
