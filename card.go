package main

import (
	"fmt"
	"github.com/nicolas-nannoni/fingy-go-client/events"
)

func RegisterCard(card CardSwipe) {

	var evt events.Event = events.Event{Path: fmt.Sprintf("alfred.accessRequest.card.%s", card.id)}
	events.Notify(evt)
}
