package main

import (
	log "github.com/Sirupsen/logrus"
	fingy "github.com/nicolas-nannoni/fingy-go-client"
)

const (
	openOrder  = "==OPEN=="
	closeOrder = "==CLOSE=="
)

func openDoor(c *fingy.Context) {
	log.Infof("Opening the door!")
	sendChannel <- openOrder
}

func closeDoor(c *fingy.Context) {
	log.Infof("Closing the door!")
	sendChannel <- closeOrder
}
