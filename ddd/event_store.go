package ddd

import (
	"encoding/json"
	"sync"
)

// EventStore implements es.AggregateRootStore
type EventStore struct {
	events map[string]eventMsg
	mu     sync.Mutex
}

type eventMsg struct {
	eventName string
	event     json.RawMessage
}

var _ AggregateRootStore = (*EventStore)(nil)

func (e EventStore) Load(ctx interface{}, aggregate *AggregateRoot) error {
	//TODO implement me
	panic("implement me")
}

func (e EventStore) Save(ctx interface{}, aggregate *AggregateRoot) error {
	//TODO implement me
	panic("implement me")
}
