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
	"github.com/go-playground/validator/v10"
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
		log.Printf("error converting event to json: %s", msg)  
		log.Printf("detail: %#v", err)
		return
	}  

	conn := database.NewDatabaseConn[models.OrderEvent]()
	isProcessed, err := conn.EventAlreadyProcess(orderEvent);    

	if err != nil {
		log.Printf("error %#v",err)
		errorEvent := models.NewErrorEvent(orderEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error unable to verify is event if processed")   
		return
	}  

	if isProcessed {
		log.Printf("event id: %v, already processed", orderEvent.EventBase.EventId )
		return
	}

	validator := validator.New() 
	validator.SetTagName("binding")
	err = validator.Struct(orderEvent) 
	if err != nil {
		errorEvent := models.NewErrorEvent(orderEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error validating error event to json: %+v", err)  
		return
	}

	err = conn.InsertProcessedEvent(orderEvent); 

	if err != nil {
		errorEvent := models.NewErrorEvent(orderEvent); 
		producer.PublishEvent(errorEvent);
		log.Fatal("unable to register order into database") 
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

	log.Printf("Order id: %s received",orderEvent.EventBase.EventId); 
	log.Printf("Order details: %#v", orderEvent)
}
