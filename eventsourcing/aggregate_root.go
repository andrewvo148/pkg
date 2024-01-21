package eventsourcing

type AggregateRoot interface {
	ID() string
	Version() int
}
