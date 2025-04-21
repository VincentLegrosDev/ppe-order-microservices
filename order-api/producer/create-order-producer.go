package producer

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"ppe4peeps.com/order-api/models"
)

func PublishOrderEvent(orderEvent models.OrderEvent) error {

	orderInBytes, err := json.Marshal(orderEvent)
	if err != nil {
		return err
	}

	producer, err := newProducer()

	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: orderEvent.Topic(),
		Value: sarama.StringEncoder(orderInBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
    log.Printf("error: (%s)", err)
		return err
	}

	log.Printf("Order is stored in topic(%s)/partition(%d)/offset(%d)\n", orderEvent.Topic(), partition, offset)

	return nil
}
