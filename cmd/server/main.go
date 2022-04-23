package main

import (
	"fmt"
	"net/http"
	"os"

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

	e.Static("/audio", "static")

	e.GET("/", controller.SayHelloWorld)

	e.GET("/stream", func(c echo.Context) error {
		f, err := os.Open("../../static/4.mp3")

		if err != nil {
			return err
		}
		return c.Stream(http.StatusOK, "audio/mpeg", f)
	})

	e.Logger.Fatal(e.Start(port))
}
