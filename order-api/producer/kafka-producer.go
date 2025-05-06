package producer

import (
	"github.com/IBM/sarama"
  "os"
)

func newProducer() (sarama.SyncProducer, error) { 

	kafkaServerAddress := os.Getenv("KAFKA_SERVER_URL")  

	brokers := []string{kafkaServerAddress}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1
	return sarama.NewSyncProducer(brokers, config)
}
