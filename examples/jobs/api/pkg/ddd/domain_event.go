package ddd


type DomainEvent interface {
	AggregateID() string
	// EventType() string
	// EventVersion() string
}