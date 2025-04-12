package services

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"ppe4peeps.com/order-api/models"
)

const topic = "OrderReceived"

type ReceivedOrderEventBody struct {
	Id       string           `json:"id"`
	Products []models.Product `json:"products"`
	Customer models.Customer  `json:"customer"`
}

type OrderReceivedEvent struct {
	EventBase models.EventBase       `json:"eventBase"`
	EventBody ReceivedOrderEventBody `json:"eventBody"`
}

func ProduceOrderEvent(createOrderEvent OrderReceivedEvent) error {

	orderInBytes, err := json.Marshal(createOrderEvent)
	if err != nil {
		return err
	}

	producer, err := newProducer()

	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(orderInBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Order is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
