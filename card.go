package main

import (
	"fmt"
	"github.com/nicolas-nannoni/fingy-gateway/events"
	fingy "github.com/nicolas-nannoni/fingy-go-client"
)

func RegisterCard(card CardSwipe) {

	var evt events.Event = events.Event{Path: fmt.Sprintf("/card/%s/swiped", card.id)}
	fingy.F.Send(evt)
}
