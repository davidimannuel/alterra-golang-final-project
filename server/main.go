package main

import (
	"log"

	"alterra-golang-final-project/businesses"
	"alterra-golang-final-project/configs"
	"alterra-golang-final-project/server/bootstraps"

	"github.com/labstack/echo/v4"
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

	ctxUc := businesses.ContextUC{
		AppHost: configs.AppHost,
		DB:      configs.DB,
	}

	//init router
	e := echo.New()
	boot := bootstraps.Bootstrap{
		App:       e,
		ContextUC: ctxUc,
	}
	boot.RegisterRoute()

	e.Logger.Fatal(boot.App.Start(configs.AppHost))
}
