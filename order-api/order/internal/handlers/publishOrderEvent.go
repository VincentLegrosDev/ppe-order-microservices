package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ppe4peeps.com/order-api/models"
	"ppe4peeps.com/order-api/producer"
)

func createOrderEvent(topic models.TopicName, order models.Order) (models.OrderEvent, error) {
	switch topic {

	case models.OrderReceived:
		return models.NewOrderReceivedEvent(order), nil
	case models.OrderConfirmed:
		return models.NewOrderConfirmedEvent(order), nil
	case models.OrderPackedAndPicked:
		return models.NewOrderPackedAndPickedEvent(order), nil
	default:
		return models.OrderEvent{}, errors.New("Unvalid order topic")
	}
}

func PublishOrderEvent(c *gin.Context, topic models.TopicName) error {
	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		log.Printf("Error serializing post data (%s)\n", err)
		return err
	}

	newOrderEvent, err := createOrderEvent(topic, order)

	if err != nil {
		log.Printf("Error creating event (%s)\n", err)
		return err
	}

	producer.PublishOrderEvent(newOrderEvent)

	log.Printf("Success (%s)\n", err)

	c.String(http.StatusOK, "ok")

	return nil
}
