package main

import (
	"fmt"
	"net/http"

	"github.com/teezzan/padio/process"
)

var Port = fmt.Sprintf(":%d", 3000)

func main() {
	fmt.Println("Getting Started")
	go process.Init()

	http.HandleFunc("/", HelloHandler)

	fmt.Println("Listening...")
	// log.Fatal(http.ListenAndServe(Port, nil))
	queue := &process.Queue
	for {
		buff := queue.BufferValue()
		if buff != nil {
			fmt.Println("0 :", buff[0])
			fmt.Println("1 :", buff[1])
		}
	}
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")
}
