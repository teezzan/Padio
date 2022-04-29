package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teezzan/padio/controller"
	"github.com/teezzan/padio/process"
)

var Port = fmt.Sprintf(":%d", 3000)

func main() {
	fmt.Println("Getting Started")
	go process.Init()

	http.HandleFunc("/", controller.AudioHandler)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(Port, nil))

}
