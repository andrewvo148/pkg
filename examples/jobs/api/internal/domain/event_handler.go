package domain

import "context"

type (
	EventHandler interface {
		Handle(ctx context.Context, de DomainEvent) error
	}

	eventHandler struct {
		h EventHandler
	}

	EventDispatcher struct {
		handlers []event
	}
)
