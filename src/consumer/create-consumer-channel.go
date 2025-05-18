package consumer

import (
	"github.com/IBM/sarama"
	"os"
	"os/signal"
	"syscall"
	"log"
	"ppe4peeps.com/services/topics"
)

func NewConsumerChannel(topic topics.TopicName, processMessage func (msg string))  { 
	
	worker, err := NewConsumer() 

	if err != nil {
		panic(err)
	} 

	consumer, err := worker.ConsumePartition(topic.String(), 0, sarama.OffsetOldest)

	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)


	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
    		log.Printf("error: (%s)", err)
			case msg := <-consumer.Messages(): 
				processMessage(string(msg.Value)) 
			case <- sigchan:
				log.Printf("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}() 

	<-doneCh

	if err := worker.Close(); err != nil{
		panic(err)
	} 
}
