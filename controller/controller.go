package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func SayHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func StreamAudio(c echo.Context) error {
	// c.Response().Header().Set(echo.HeaderContentType, "audio/mpeg")
	// c.Response().WriteHeader(http.StatusOK)

	// f, err := os.Open("../static/2.mp3")
	// if err != nil {
	// 	return err
	// }
	// return c.Stream(http.StatusOK, "audio/mpeg", f)

	f, err := os.Open("../static/cores.png")
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "image/png", f)
}
