package domain

import (
	"context"
	"sync"
)

type (
	Event struct {
		Name    string
		Payload []string
	}

	EventHandler interface {
		Handle(ctx context.Context, event Event) error
	}

	//eventHandler func(ctx context.Context, event Event) error

	EventDispatcher struct {
		m   map[string][]EventHandler
		mut sync.RWMutex

	}
)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		m: make(map[string][]EventHandler),
	}
}

func (ed *EventDispatcher) Publish(ctx context.Context, event Event) error {
	ed.mut.RLock()
	defer ed.mut.RUnlock()

	if handlers, exists := ed.m[event.Name]; exists {
		for _, h := range handlers {
			err := h.Handle(ctx, event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ed *EventDispatcher) Subscribe(h EventHandler, eventName string) error {
	ed.mut.Lock()
	defer ed.mut.Unlock()

	if handlers, exists := ed.m[eventName]; exists {
		ed.m[eventName] = append(handlers, h)
	} else {
		ed.m[eventName] = []EventHandler{h}

	}

	return nil
}

