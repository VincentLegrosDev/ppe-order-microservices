package main

import (
	"github.com/google/uuid"
	"ppe4peeps.com/order-api/models"
	"ppe4peeps.com/order-api/producer"
	"time"
)

func main() {

	product := models.Product{Quantity: 2, ProductId: uuid.New()}

	var order = models.Order{
		OrderId:  uuid.New(),
		Products: []models.Product{product},
		Customer: models.Customer{
			FirstName:     "Jean",
			LastName:      "Jacques",
			EmailAddresse: "emia@email.com",
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
			EventName:      models.OrderReceived,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}

	producer.PublishOrderEvent(orderReceivedEvent)

}
