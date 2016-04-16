package main

import (
	log "github.com/Sirupsen/logrus"
	fingy "github.com/nicolas-nannoni/fingy-go-client"
	"time"
)

func main() {

	setupAndConnectToFingy()

	for {
		t := time.NewTimer(time.Second * 5)
		<-t.C

		RegisterCard(CardSwipe{id: "card-123"})
	}

	select {}
}

// Setup FingyClient and initiate the connection to Fingy
func setupAndConnectToFingy() {

	fingy.F.ServiceId = "alfred"
	fingy.F.DeviceId = "abcdef1234"
	registerEvents()

	go fingy.F.Begin()
}

// Register supported events, received from Alfred server via the Fingy connection
func registerEvents() {
	fingy.F.Router.Entry("/door/open", openDoor)
}

// Toggle the door GPIO pin for some time to open it
func openDoor(c *fingy.Context) {
	log.Info("Opening door")
}
