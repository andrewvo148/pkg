package ddd

type Aggregate interface {
	Entity
	ApplyEvent(command Command) error
	ApplySnapshot(snapshot Snapshot) error
}
