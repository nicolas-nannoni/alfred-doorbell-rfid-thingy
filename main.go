package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	fingy "github.com/nicolas-nannoni/fingy-go-client"
)

var (
	deviceId  = flag.String("deviceId", "571a8511d4c6c5edf1488853", "the Fingy deviceId")
	serviceId = flag.String("serviceId", "alfred", "the Fingy serviceId")
	fingyHost = flag.String("fingyHost", "localhost:8080", "the Fingy Gateway Hostname & Port")
	debug     = flag.Bool("debug", false, "enable debug logging")
)

func main() {

	flag.Parse()

	setupLogging()
	setupAndConnectToFingy()
	go SerialLoop()

	select {}
}

func setupLogging() {
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}

// Setup FingyClient and initiate the connection to Fingy
func setupAndConnectToFingy() {

	fingy.F.ServiceId = *serviceId
	fingy.F.DeviceId = *deviceId
	fingy.F.FingyHost = *fingyHost
	registerEvents()

	go fingy.F.Begin()
}

// Register supported events, received from Alfred server via the Fingy connection
func registerEvents() {
	fingy.F.Router.Entry("/open", openDoor)
	fingy.F.Router.Entry("/close", closeDoor)
}
