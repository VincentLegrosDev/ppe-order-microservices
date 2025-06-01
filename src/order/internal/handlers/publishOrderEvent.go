package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"ppe4peeps.com/services/models"
	"ppe4peeps.com/services/producer"
	"ppe4peeps.com/services/topics"
)

func createOrderEvent(topic topics.TopicName, order models.Order) (models.OrderEvent, error) {
	switch topic {

	case topics.OrderReceived:
		return models.NewOrderReceivedEvent(order), nil
	case topics.OrderConfirmed:
		return models.NewOrderConfirmedEvent(order), nil
	case topics.OrderPackedAndPicked:
		return models.NewOrderPackedAndPickedEvent(order), nil
	default:
		return models.OrderEvent{}, errors.New("Unvalid order topic")
	}
}

func PublishOrderEvent(c *gin.Context, topic topics.TopicName) error {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		log.Printf("Error serializing post data (%s)\n", err) 
		errorEvent := models.NewErrorEvent(models.NewOrderReceivedEvent(order))
		producer.PublishEvent(errorEvent) 
		c.AbortWithError(http.StatusBadRequest, err);
		return err
	}

	newOrderEvent, err := createOrderEvent(topic, order)

	if err != nil {
		log.Printf("Error creating event (%s)\n", err)
		c.AbortWithError(http.StatusBadRequest, err);
		errorEvent := models.NewErrorEvent(models.NewOrderReceivedEvent(order))
		producer.PublishEvent(errorEvent) 
		return err
	}

	producer.PublishEvent(newOrderEvent)
	log.Printf("Success (%s)\n", newOrderEvent.Topic())
	c.String(http.StatusOK, "ok") 

	if err := publishOrderCountMetric(order); err != nil {
		return err
	} 


	return nil
}


func publishOrderCountMetric(o models.Order) error {

	tag := models.Tag{
		Name: "new_products_order",
		Value: strconv.Itoa(len(o.Products)), 
	} 

	if err := producer.PublishEvent(models.NewOrderCountEvent( models.OrderCountMetric { 
		Count: 1, 
		Tags:  [] models.Tag{tag},
	})); err != nil {

		return err 
	}

	return nil
}
