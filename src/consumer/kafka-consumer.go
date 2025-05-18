package consumer

import (
	"github.com/IBM/sarama"
	"os"
	"log"
) 

func NewConsumer() (sarama.Consumer, error) {

	kafkaServerAddress := os.Getenv("KAFKA_SERVER_URL")   

	log.Printf("start consumer kafkaServerAddress: %s", kafkaServerAddress)
	brokers := []string{kafkaServerAddress}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Net.MaxOpenRequests = 1 

	return sarama.NewConsumer(brokers, config)
}
