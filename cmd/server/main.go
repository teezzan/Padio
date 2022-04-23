package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/teezzan/padio/config"
	"github.com/teezzan/padio/controller"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	port := fmt.Sprintf(":%d", config.HTTP.Port)
	e := echo.New()

	e.GET("/", controller.SayHelloWorld)

	e.Logger.Fatal(e.Start(port))
}
