package models

import (
	"github.com/google/uuid"
	"time"
)

type EventName string

type TopicName string

const (
	OrderReceived        TopicName = "OrderReceived"
	OrderConfirmed       TopicName = "OrderConfirmed"
	OrderPackedAndPicked TopicName = "OrderPackedAndPicked"
)

type EventBase struct {
	EventId        uuid.UUID
	EventName      TopicName
	EventTimestamp time.Time
}

type OrderEvent struct {
	EventBase EventBase
	EventBody Order
}

func (event OrderEvent) Topic() string {
	return string(event.EventBase.EventName)
}

func NewOrderReceivedEvent(order Order) OrderEvent {
	return OrderEvent{
		EventBase: EventBase{
			EventId:        uuid.New(),
			EventName:      OrderReceived,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}
}

func NewOrderConfirmedEvent(order Order) OrderEvent {
	return OrderEvent{
		EventBase: EventBase{
			EventId:        uuid.New(),
			EventName:      OrderConfirmed,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}
}

func NewOrderPackedAndPickedEvent(order Order) OrderEvent {
	return OrderEvent{
		EventBase: EventBase{
			EventId:        uuid.New(),
			EventName:      OrderPackedAndPicked,
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}
}
