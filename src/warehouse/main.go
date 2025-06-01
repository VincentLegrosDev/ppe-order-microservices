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
	"strconv"
)

func main() { 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumer.NewConsumerChannel(topics.OrderConfirmed, consumeOrderConfirmedEvent) 
}

func consumeOrderConfirmedEvent(msg string) {
	var orderConfirmedEvent models.OrderEvent
	err := json.Unmarshal([]byte(msg), &orderConfirmedEvent)

	if err != nil {
		errorEvent := models.NewErrorEvent(orderConfirmedEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error converting event to json: %s", msg)  
		log.Printf("detail: %#v", err)
		return
	} 

	conn := database.NewDatabaseConn[models.OrderEvent]()
	isProcessed, err := conn.EventAlreadyProcess(orderConfirmedEvent);    

	if err != nil {
		errorEvent := models.NewErrorEvent(orderConfirmedEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error unable to verify is event if processed")   
		return
	}  
	
	if isProcessed {
		log.Printf("event id: %v, already processed", orderConfirmedEvent.EventBase.EventId )
		return
	}

	validator := validator.New() 
	validator.SetTagName("binding")
	err = validator.Struct(orderConfirmedEvent) 
	if err != nil {
		errorEvent := models.NewErrorEvent(orderConfirmedEvent) 
		producer.PublishEvent(errorEvent);
		log.Printf("error validating event to json: %+v", err)  
		return
	}


	err = conn.InsertProcessedEvent(orderConfirmedEvent); 

	if err != nil {
		errorEvent := models.NewErrorEvent(orderConfirmedEvent) 
		producer.PublishEvent(errorEvent);
		log.Fatal("unable to register order into database") 
		return
	}

	notificationEvent := models.NewNotificationEvent(models.Notification {
		Type: "OrderConfirmed",
		Recipient: orderConfirmedEvent.Order().Customer.EmailAddress,
		From: "orders@ppe4peeps.com",
		Subject: "Order confirmed", 
		Body: "Your order is been confirmed and will be shipped shortly, thank you to do business with us", 
	})
	err = producer.PublishEvent(notificationEvent);


	if err != nil { 
		errorEvent := models.NewErrorEvent(notificationEvent); 
		producer.PublishEvent(errorEvent);
		log.Fatal("error sending notification") 
		return
	}

	log.Printf("Order id: %v is confirmed",orderConfirmedEvent.EventBase.EventId); 
	log.Printf("order id: %v notification send ", orderConfirmedEvent.EventBase.EventId)

	publishOrderTimeMetric(orderConfirmedEvent.Order());
}



func publishOrderTimeMetric(order models.Order) error { 

		tag1 := models.Tag{
			Name:  "products_ordered",
			Value: strconv.Itoa(len(order.Products)),
		}
		tag2 := models.Tag{
			Name:  "order_id",
			Value: order.OrderId.String(),
		}
		tag3 :=  models.Tag{
			Name:  "process_step",
			Value: "warehouse",
		}
		tags := []models.Tag{tag1, tag2, tag3}

		if err := producer.PublishEvent(models.NewOrderTimeEvent(models.OrderTimeMetric{
			Tags: tags,
			Count: 1,

		})); err != nil {
			return err;
		}

	return  nil 
}
