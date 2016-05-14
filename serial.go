package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/tarm/serial"
	"regexp"
	"time"
)

var (
	portName = flag.String("serialPort", "/dev/ttyUSB0", "the Arduino Serial port")

	cardRe = regexp.MustCompile(`\*CARD\*:([\d\w]*)`)
	keyRe  = regexp.MustCompile(`\*KEY\*:(\d)`)
	bellRe = regexp.MustCompile(`\*BELL\*`)

	sendChannel = make(chan string)
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

func SerialLoop() {

	port := openPort()

	log.Printf("Port %s opened", *portName)

	for {
		select {
		case msg := <-sendChannel:
			send(port, msg)
		default:
			listen(port)
		}
	}

}

func openPort() (port *serial.Port) {

	c := &serial.Config{Name: *portName, Baud: 9600, ReadTimeout: time.Second * 1}

	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func closePort(port *serial.Port) {
	log.Infof("Closing port %s", *portName)
	port.Close()
}

func send(port *serial.Port, msg string) {
	log.Debugf("Message to send: %s", msg)
	port.Write([]byte(msg + "\n"))
}

func listen(p *serial.Port) {
	in, ok := readBytes(p)
	if !ok {
		return
	}

	evt, err := parseInput(in)
	if err != nil {
		log.Println(err)
		return
	}

	err = dispatchEvent(evt)
	if err != nil {
		log.Println(err)
		return
	}
}

func readBytes(p *serial.Port) (str string, ok bool) {

	in := make([]byte, 32)
	n, err := p.Read(in)
	if n == 0 {
		log.Debugf("No data on serial: %v", n)
		ok = false
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	str = string(in[:n])
	log.Debugf("Received input: %s", str)

	ok = true
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
