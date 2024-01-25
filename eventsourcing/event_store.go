package eventsourcing

type EventStore interface {
	Publish(events ...Event) error
}
