package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teezzan/padio/process"
)

var Port = fmt.Sprintf(":%d", 3000)

func main() {
	fmt.Println("Getting Started")
	go process.Init()

	http.HandleFunc("/", HelloHandler)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(Port, nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")
}
