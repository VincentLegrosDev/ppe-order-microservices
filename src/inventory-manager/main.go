package main 

import (
	"ppe4peeps.com/services/consumer"
	"ppe4peeps.com/services/models"
	"ppe4peeps.com/services/topics"
	"ppe4peeps.com/services/producer"
	"ppe4peeps.com/services/database"
	"log"
	"github.com/joho/godotenv"
	"encoding/json"
)

func main() { 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumer.NewConsumerChannel(topics.OrderReceived, consumeOrderReceivedEvent)
}

func consumeOrderReceivedEvent(msg string) {
	var orderEvent models.OrderEvent
	err := json.Unmarshal([]byte(msg), &orderEvent)

	if err != nil {
		errorEvent := models.NewErrorEvent(orderEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error confirming order")  
	} 

	conn := database.NewDatabaseConn()

	isProcessed, err := conn.OrderAlreadyProcess(orderEvent);    

	if err != nil {
		errorEvent := models.NewErrorEvent(orderEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error unable to verify is event if processed")   
	}  
	
	if isProcessed {
		log.Printf("event id: %v, already processed", orderEvent.EventBase.EventId )
		return
	}

	orderConfirmedEvent := models.NewOrderConfirmedEvent(orderEvent.Order())
	err = producer.PublishEvent(orderConfirmedEvent) 


	if err != nil { 
		errorEvent := models.NewErrorEvent(orderConfirmedEvent); 
		producer.PublishEvent(errorEvent);
		log.Fatal("error confirming order") 
		return
	}

	err = conn.InsertProcessedOrder(orderEvent); 

	if err != nil {
		log.Fatal("unable to register order into database") 
		return
	}

	log.Printf("Order id: %s processed",orderEvent.EventBase.EventId); 
	log.Printf("Order details: %#v", orderEvent)
}
