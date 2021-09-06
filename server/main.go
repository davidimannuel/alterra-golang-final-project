package main

import (
	"fmt"
	"log"

	"alterra-golang-final-project/config"
	"alterra-golang-final-project/server/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// read config
	configs, err := config.LoadConfigs()
	if err != nil {
		log.Fatal("Error load config file")
	}
	fmt.Println("configs", configs)
	defer func() {
		log.Print("defer function")
	}()

	//init router
	e := echo.New()
	r := routes.Router{
		Echo:    e,
		Configs: configs,
	}
	r.RegisterRoute()

	e.Logger.Fatal(r.Echo.Start(configs.AppHost))
}
