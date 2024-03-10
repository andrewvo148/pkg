package ddd

type AggregateRoot struct {
	domainEvents []DomainEvent
}

func (ar *AggregateRoot) AddEvent(e DomainEvent) {
	ar.domainEvents = append(ar.domainEvents, e)
}

func (ar *AggregateRoot) ClearEvents() {
	ar.domainEvents = []DomainEvent{}
}
