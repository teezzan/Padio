package controller

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tosone/minimp3"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")
}

func AudioHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	var file *os.File
	if file, err = os.Open("static/04. Unaaji.mp3"); err != nil {
		log.Fatal(err)
	}

	var dec *minimp3.Decoder
	if dec, err = minimp3.NewDecoder(file); err != nil {
		log.Fatal(err)
	}
	started := dec.Started()
	<-started

	log.Printf("Convert audio sample rate: %d, channels: %d\n", dec.SampleRate, dec.Channels)

	// for {
	// var data = make([]byte, 1024)
	// 	_, err := dec.Read(data)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		break
	// 	}
	// 	// player.Write(data)

	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	// 	w.Header().Set("Connection", "Keep-Alive")
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("X-Content-Type-Options", "nosniff")
	// 	fmt.Fprint(w, data)
	// 	flusher.Flush() // Trigger "chunked" encoding and send a chunk...
	// 	fmt.Println("Sending: ", "data")
	// 	time.Sleep(500 * time.Millisecond)
	// }
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Content-Type", "audio/mpeg")
	// Don't need these this bc manually flushing sets this header
	// w.Header().Set("Transfer-Encoding", "chunked")

	// make sure this header is set
	w.Header().Set("X-Content-Type-Options", "nosniff")

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		var data = make([]byte, 1024)
		_, err := dec.Read(data)
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		for t := range ticker.C {
			// #2 add '\n'
			binary.Write(w, binary.BigEndian, data)

			fmt.Println("Tick at", t)
			flusher.Flush()
		}
	}()
	time.Sleep(time.Second * 50)
	ticker.Stop()
}

// queue := &process.Queue
// 	for {
// 		buff := queue.BufferValue()
// 		if buff != nil {
// 			fmt.Println("0 :", buff[0])
// 			fmt.Println("1 :", buff[1])
// 		}
// 	}
