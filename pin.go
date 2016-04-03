package main

import (
	"log"
	"strconv"
	"time"
)

const (
	pinInputLength = 6
	pinTimeout     = 2
)

var (
	pinInput = make([]int, pinInputLength)
	pinIndex = 0

	pinTimer = time.NewTimer(time.Second * pinTimeout)
)

type KeyStroke struct {
	InputEvent
	value string
}

func init() {
	ResetPin()
}

func RegisterStroke(key KeyStroke) {

	i, err := strconv.Atoi(key.value)
	if err != nil {
		log.Printf("Invalid PIN digit received: %s", key.value)
		return
	}

	storeDigit(i)
}

func ResetPin() {
	pinInput = make([]int, pinInputLength)
	pinIndex = 0
}

func flushPin() {
	log.Printf("Flushing PIN! %v", pinInput)
	pinTimer.Stop()
	ResetPin()
}

func restartTimeout() {

	pinTimer.Stop()
	pinTimer = time.NewTimer(time.Second * pinTimeout)

	go func() {
		<-pinTimer.C
		flushPin()
	}()
}

func storeDigit(i int) {
	pinInput[pinIndex] = i
	pinIndex++
	if pinIndex > len(pinInput)-1 {
		flushPin()
		return
	}
	restartTimeout()
}
