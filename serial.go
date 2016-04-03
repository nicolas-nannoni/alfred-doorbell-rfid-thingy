package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"regexp"
)

const (
	portName = "/dev/tty.wchusbserial1420"
)

var (
	cardRe = regexp.MustCompile(`\*CARD\*:([\d\w]*)`)
	keyRe  = regexp.MustCompile(`\*KEY\*:(\d)`)
	bellRe = regexp.MustCompile(`\*BELL\*`)
)

type InputEvent interface {
}

type CardSwipe struct {
	InputEvent
	id string
}

type DoorBell struct {
	InputEvent
}

func ListenSerial() {

	port := openPort()

	log.Printf("Port %s opened", portName)

	for {
		in := listen(port)

		evt, err := parseInput(in)
		if err != nil {
			log.Println(err)
			continue
		}

		err = dispatchEvent(evt)
		if err != nil {
			log.Println(err)
			continue
		}
	}

}

func openPort() (port *serial.Port) {

	c := &serial.Config{Name: portName, Baud: 9600}

	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func closePort(port *serial.Port) {
	log.Printf("Closing port %s", portName)
	port.Close()
}

func listen(p *serial.Port) (str string) {

	in := make([]byte, 32)
	n, err := p.Read(in)
	if err != nil {
		log.Fatal(err)
	}

	str = string(in[:n])
	log.Printf("Received input: %s", str)

	return
}

func parseInput(in string) (evt InputEvent, err error) {

	switch {
	case cardRe.MatchString(in):
		card := CardSwipe{id: cardRe.FindStringSubmatch(in)[1]}
		log.Printf("Card swipped %s", card.id)
		return card, err
	case keyRe.MatchString(in):
		key := KeyStroke{value: keyRe.FindStringSubmatch(in)[1]}
		log.Printf("Pressed key: %s", key.value)
		return key, err
	case bellRe.MatchString(in):
		log.Println("Door bell")
		return DoorBell{}, err
	default:
		err = fmt.Errorf("Unknown input received: %q", in)
	}
	return
}

func dispatchEvent(evt InputEvent) (err error) {

	switch evt.(type) {
	case CardSwipe:
		RegisterCard(evt.(CardSwipe))
	case KeyStroke:
		RegisterStroke(evt.(KeyStroke))
	case DoorBell:
		OnDoorBell(evt.(DoorBell))
	default:
		err = fmt.Errorf("Unsupported InputEvent received: %T", evt)
	}
	return
}
