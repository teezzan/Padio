package controller

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"time"

	"github.com/teezzan/padio/process"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")
}

func AudioHandler(w http.ResponseWriter, r *http.Request) {
	queue := &process.Queue

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {

		for t := range ticker.C {
			buff := queue.BufferValue()
			binary.Write(w, binary.BigEndian, buff)

			fmt.Println("Tick at", t)
			flusher.Flush()
		}
	}()
	time.Sleep(time.Second * 50)
	ticker.Stop()
}
