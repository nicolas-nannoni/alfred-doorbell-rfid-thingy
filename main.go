package main

import (
	fingy "github.com/nicolas-nannoni/fingy-go-client"
	"log"
)

func main() {

	go ListenSerial()

	fingy.Router.Entry("/hello", func(c *fingy.Context) {
		log.Print("hello")
	})

	fingy.DeviceId = "abcdef1234"

	go fingy.Run()

	select {}
}
