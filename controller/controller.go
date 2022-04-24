package controller

import (
	"encoding/binary"
	"net/http"

	"github.com/labstack/echo/v4"
)

const sampleRate = 44100
const seconds = 2

var buffer = make([]float32, sampleRate*seconds)

func SayHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func StreamAudio(c echo.Context) {
	w := c.Response().Writer
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/wave")
	for {
		binary.Write(w, binary.BigEndian, &buffer)
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...

	}
}
