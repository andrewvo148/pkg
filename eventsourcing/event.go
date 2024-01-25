package eventsourcing

import "time"

type IDer interface {
	ID() string
}

type EventPayload interface{}
type Event interface {
	IDer
	EventName() string
	Payload() EventPayload
	OccurredAt() time.Time
}

type event struct {
	id         string
	eventName  string
	payload    EventPayload
	occurredAt time.Time
}

var _ Event = (*event)(nil)

func NewEvent(id, name string, payload EventPayload) event {
	return event{
		id:         id,
		payload:    payload,
		occurredAt: time.Now(),
	}
}

func (e event) ID() string {
	return e.id
}

func (e event) EventName() string {
	return e.eventName
}

func (e event) Payload() EventPayload {
	return e.payload
}

func (e event) OccurredAt() time.Time {
	return e.occurredAt
}
