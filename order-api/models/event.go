package models

import (
	"time"
	"github.com/google/uuid"
)

type EventBase struct {
	EventId        uuid.UUID
	EventName      string
	EventTimestamp time.Time
}

type OrderReceivedEvent struct {
	EventBase EventBase
	EventBody Order
}
