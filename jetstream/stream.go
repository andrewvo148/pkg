package jetstream

import (
	"context"
	"github.com/andrewvo148/pkg/log"
	"github.com/nats-io/nats.go"
)

const maxRetries = 5

type Stream struct {
	logger log.Logger
	js     nats.JetStream
}

func (s *Stream) Publish(ctx context.Context, topicName string) (err error) {
	var p nats.PubAck
	p, err = s.js.PublishMsgAsync(&nats.Msg{
		Subject: "",
		Reply:   "",
		Header:  nil,
		Data:    nil,
		Sub:     nil,
	}, nats.MsgId())
	if err != nil {
		return
	}

	// retry a handful of times to publish the messages
	go func(future nats.PubAckFuture, tries int) {
		var err error

		for {
			select {
			case <-future.Ok(): // publish acknowledged
				return
			case <-future.Err(): // error ignored; try again
				// TODO add some variable delay between tries
				tries -= 1
				if tries <= 0 {
					return
				}

				future, err = s.js.PublishMsgAsync(future.Msg())
				if err != nil {
					// TODO do more than give up
					return
					nats.ErrMsgNotFound
				}
			}
		}
	}(p, maxRetries)
	return
}
