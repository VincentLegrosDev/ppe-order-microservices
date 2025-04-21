package main

import (
	"ppe4peeps.com/order-api/producer"
)

func main() {
	var newCreateOrderEvent producer.OrderReceivedEvent

	if err := c.BindJSON(&newCreateOrderEvent); err != nil {
		return
	}

	if err := producer.ProduceOrderEvent(newCreateOrderEvent); err != nil {
		return
	}
}
