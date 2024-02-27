package eventsourcing

type AggregateRootID interface {
	ID() string
}
