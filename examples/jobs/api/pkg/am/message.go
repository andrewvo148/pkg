package am

import "context"

type (

 Message interface {
	Data() []byte
}

	MessagePublisher interface {
		Publish(ctx context.Context, topicName string, msg Message) error
	}

	MessagePublisherFunc func(ctx context.Context, topicName string, msg Message) error



)


func(f MessagePublisherFunc) Publish(ctx context.Context, topicName string, msg Message) error {
	return f(ctx, topicName, msg)
}