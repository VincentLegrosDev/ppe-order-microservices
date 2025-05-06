package models

import (
	"github.com/google/uuid"
	"ppe4peeps.com/services/topics"
	"time"
)

type Event interface {
//	OrderEvent | ErrorEvent | NotificationEvent
	Topic() string
}

type EventName string

type EventBase struct {
	EventId        uuid.UUID
	EventName      topics.TopicName
	EventTimestamp time.Time
}

type OrderEvent struct{
	EventBase EventBase
	EventBody Order
}  

type ErrorEvent struct{
	EventBase EventBase
	EventBody OrderEvent
}

type NotificationEvent struct {
	EventBase EventBase
	EventBody Notification
}

func (event OrderEvent) Topic() string {
	return string(event.EventBase.EventName)
}  

func (event NotificationEvent) Topic() string {
	return string(event.EventBase.EventName)
}  

func (event ErrorEvent) Topic() string {
	return string(event.EventBase.EventName)
}


func (event OrderEvent) Order() Order  {
	return event.EventBody
}

func NewOrderReceivedEvent(order Order) OrderEvent {
	return OrderEvent{
		EventBase: EventBase{
			EventId:        uuid.New(),
			EventName:      topics.OrderReceived,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}
}

func NewOrderConfirmedEvent(order Order) OrderEvent {
	return OrderEvent{
		EventBase: EventBase{
			EventId:        uuid.New(),
			EventName:      topics.OrderConfirmed,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}
}

func NewOrderPackedAndPickedEvent(order Order) OrderEvent {
	return OrderEvent{
		EventBase: EventBase{
			EventId:        uuid.New(),
			EventName:      topics.OrderPackedAndPicked,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	} 
}

func NewErrorEvent(event OrderEvent) ErrorEvent {  
	return ErrorEvent{
		EventBase: EventBase{
			EventId: uuid.New(),
			EventName: topics.DeadQueueLetter,
			EventTimestamp: time.Now(),
		}, 
		EventBody: event,
	}
} 

func NewNotificationEvent(notification Notification) NotificationEvent {
	return NotificationEvent{
		EventBase: EventBase {
			EventId: uuid.New(), 
			EventName: topics.Notification, 
			EventTimestamp: time.Now(), 
		}, 
		EventBody: notification,
	}

}


