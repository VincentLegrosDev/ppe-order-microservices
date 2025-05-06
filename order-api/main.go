package main

import (
	"github.com/google/uuid"
	"ppe4peeps.com/services/models"
	"ppe4peeps.com/services/producer"
	"ppe4peeps.com/services/topics"
	"time"
)

func main() {

	product := models.Product{Quantity: 2, ProductCode: "1231",}

	var order = models.Order{
		OrderId:  uuid.New(),
		Products: []models.Product{product},
		Customer: models.Customer{
			FirstName:     "Jean",
			LastName:      "Jacques",
			EmailAddress: "emia@email.com",
			ShippingAddress: models.ShippingAddress{
				Line1:      "12543 rue deschamps",
				City:       "MarieVille",
				State:      "QC",
				PostalCode: "JS@ 234",
			},
		},
	}

	var orderReceivedEvent = models.OrderEvent{
		EventBase: models.EventBase{
			EventId:        uuid.New(),
			EventName:      topics.OrderReceived,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}

	producer.PublishEvent(orderReceivedEvent)

}
