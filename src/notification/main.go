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

	consumer.NewConsumerChannel(topics.Notification,consumeNotificationEvent)
}

func consumeNotificationEvent(msg string) {
	var notification models.NotificationEvent
	err := json.Unmarshal([]byte(msg), &notification)

	if err != nil {
		errorEvent := models.NewErrorEvent(notification)   
		producer.PublishEvent(errorEvent);
		log.Printf("error confirming order")  
		return
	} 

	conn := database.NewDatabaseConn[models.NotificationEvent]() 

	isProcessed, err := conn.EventAlreadyProcess(notification) 


	if err != nil {
		errorEvent := models.NewErrorEvent(notification) 
		producer.PublishEvent(errorEvent);
		log.Printf("error unable to verify is event if processed")   
		return
	}  
	
	if isProcessed {
		log.Printf("event id: %v, already processed", notification.EventBase.EventId )
		return
	}

	validator := validator.New() 
	validator.SetTagName("binding")
	err = validator.Struct(notification) 
	if err != nil {
		errorEvent := models.NewErrorEvent(notification) 
		producer.PublishEvent(errorEvent);
		log.Printf("error validating event to json: %+v", err)  
		return
	}

	err = conn.InsertProcessedEvent(notification); 

	if err != nil {
		errorEvent := models.NewErrorEvent(notification) 
		producer.PublishEvent(errorEvent);
		log.Fatal("unable to register order into database") 
		return
	}

	log.Printf("Order id: %v, %s notification send",notification.EventBase.EventId, notification.EventBody.Type);  

}
