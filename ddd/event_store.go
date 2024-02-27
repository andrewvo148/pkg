package eventsourcing
// Persisted events
type EventStore interface {
	Publish(events ...Event) error
}
