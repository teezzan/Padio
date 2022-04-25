package main

import (
	"fmt"

	"github.com/teezzan/padio/process"
)

func main() {
	fmt.Println("Getting Started")
	go process.Init()
	fmt.Println("Entering loop")

	select {}
}
