package services

import (
	"github.com/IBM/sarama"
)

const (
	ProducerPort       = ":8080"
	KafkaServerAddress = "localhost:9092"
)

func newProducer() (sarama.SyncProducer, error) {
	brokers := []string{KafkaServerAddress}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	return sarama.NewSyncProducer(brokers, config)
}
