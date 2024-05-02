package service

import "time"

type EventPayload interface{}

type CreatedEvent struct {
	ID               string
	EventCreatedTime time.Time
	Payload          EventPayload
}

type EntityHandler interface {
	HandleCreatedEvent(ent *CreatedEvent)
}
