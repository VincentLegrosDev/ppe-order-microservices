package models

import (
	"github.com/google/uuid"
	"ppe4peeps.com/services/topics"
	"time"
)

type Event interface {
	Topic() string
	Id() uuid.UUID 
}

type EventName string

type EventBase struct {
	EventId        uuid.UUID `json:"eventId" binding:"required"`
	EventName      topics.TopicName `json:"eventName" binding:"required"`
	EventTimestamp time.Time `json:"eventTimestamp" binding:"required"`
}

type OrderEvent struct{
	EventBase EventBase `json:"eventBase" binding:"required"`
	EventBody Order `json:"eventBody" binding:"required"`
}  

type ErrorEvent [T Event] struct {
	EventBase EventBase `json:"eventBase" binding:"required"`
	EventBody T `json:"eventBody" binding:"required"`
}

type NotificationEvent struct {
	EventBase EventBase `json:"eventBase" binding:"required"`
	EventBody Notification `json:"eventBody" binding:"required"`
}

type OrderTimeEvent struct {
	EventBase EventBase `json:"eventBase" binding:"required"`
	EventBody OrderTimeMetric `json:"eventBody" binding:"required"`
}

type OrderCountEvent struct {
	EventBase EventBase `json:"eventBase" binding:"required"`
	EventBody OrderCountMetric `json:"eventBody" binding:"required"`
}

func (event OrderEvent) Topic() string {
	return string(event.EventBase.EventName)
}  

func (event NotificationEvent) Topic() string {
	return string(event.EventBase.EventName)
}  

func (event ErrorEvent [T]) Topic() string { 
	return string(event.EventBase.EventName)
}

func (event OrderCountEvent) Topic() string {
	return string(event.EventBase.EventName)
}

func (event OrderTimeEvent) Topic() string {
	return string(event.EventBase.EventName)
}

func (event OrderEvent) Id() uuid.UUID {
	return event.EventBase.EventId
}  

func (event NotificationEvent) Id()  uuid.UUID{
	return event.EventBase.EventId
}  

func (event ErrorEvent [T]) Id() uuid.UUID { 
	return event.EventBase.EventId 
}

func (event OrderCountEvent) Id() uuid.UUID {
	return event.EventBase.EventId
}

func (event OrderTimeEvent) Id() uuid.UUID {
	return event.EventBase.EventId
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

func NewErrorEvent[T Event](event T) ErrorEvent[T] {      
	return ErrorEvent[T]{
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

func NewOrderCountEvent(orderCountMetric OrderCountMetric) OrderCountEvent {
	return OrderCountEvent {
		EventBase: EventBase {
			EventId: uuid.New(),  
			EventName: topics.OrderCountMetric, 
		},
		EventBody: orderCountMetric, 
	}
}

func NewOrderTimeEvent(orderTimeMetric OrderTimeMetric) OrderTimeEvent {
	return OrderTimeEvent {
		EventBase: EventBase {
			EventId: uuid.New(),  
			EventName: topics.OrderTimeMetric, 
		},
		EventBody: orderTimeMetric, 
	}
}




