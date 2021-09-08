package main

import (
	"log"

	"keep-remind-app/businesses"
	"keep-remind-app/configs"
	"keep-remind-app/server/bootstraps"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// read config
	configs, err := configs.LoadConfigs()
	if err != nil {
		log.Fatal("Error load config file")
	}

	defer func() {
		log.Print("defer function")
	}()

	contextUC := businesses.ContextUC{
		AppHost:    configs.AppHost,
		AppTimeout: configs.AppTimeout,
		DB:         configs.DB,
	}

	//init router
	e := echo.New()
	e.Use(middleware.Logger())
	boot := bootstraps.Bootstrap{
		App:       e,
		ContextUC: contextUC,
	}
	boot.RegisterRoute()

	e.Logger.Fatal(boot.App.Start(configs.AppHost))
}
