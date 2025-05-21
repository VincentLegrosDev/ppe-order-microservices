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

	consumer.NewConsumerChannel(topics.OrderPackedAndPicked, consumeOrderPickedAndPackedEvent) 
}

func consumeOrderPickedAndPackedEvent(msg string) {
	var orderShippedEvent models.OrderEvent
	err := json.Unmarshal([]byte(msg), &orderShippedEvent)

	if err != nil {
		errorEvent := models.NewErrorEvent(orderShippedEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error converting event to json: %s", msg)  
		log.Printf("detail: %#v", err)
		return
	}  

	conn := database.NewDatabaseConn[models.OrderEvent]()
	isProcessed, err := conn.EventAlreadyProcess(orderShippedEvent);    

	if err != nil {
		errorEvent := models.NewErrorEvent(orderShippedEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error unable to verify is event if processed")   
		return
	}  
	
	if isProcessed {
		log.Printf("event id: %v, already processed", orderShippedEvent.EventBase.EventId )
		return
	}

	validator := validator.New() 
	validator.SetTagName("binding")
	err = validator.Struct(orderShippedEvent) 
	if err != nil {
		errorEvent := models.NewErrorEvent(orderShippedEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error validating event to json: %+v", err)  
		return
	}


	err = conn.InsertProcessedEvent(orderShippedEvent); 

	if err != nil { 
		errorEvent := models.NewErrorEvent(orderShippedEvent); 
		producer.PublishEvent(errorEvent);
		log.Fatal("error confirming order") 
		return
	}


	order := orderShippedEvent.Order();
	log.Printf("Cutomer: %ss %s", order.Customer.FirstName, order.Customer.LastName);
	log.Printf("ShippingAddresse, %s", order.Customer.ShippingAddress);

	notificationEvent := models.NewNotificationEvent(models.Notification {
		Type: "OrderShipped",
		Recipient: orderShippedEvent.Order().Customer.EmailAddress,
		From: "orders@ppe4peeps.com",
		Subject: "Order shipped", 
		Body: "Your order has been processed and shipped, thank you to do business with us", 
	})
	err = producer.PublishEvent(notificationEvent);

	if err != nil {
		notificationErrorEvent := models.NewErrorEvent(notificationEvent); 
		producer.PublishEvent(notificationErrorEvent);
		log.Fatal("unable to register order into database") 
		return
	}

	log.Printf("Order id: %s is shipped",orderShippedEvent.EventBase.EventId); 
}
