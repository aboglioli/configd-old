package infrastructure

import (
	"github.com/aboglioli/configd/pkg/events"
)

type InMemEventBus struct {
	subscriptions map[string][]events.SubscriptionFunc
}

func NewInMemEventBus() *InMemEventBus {
	return &InMemEventBus{
		subscriptions: make(map[string][]events.SubscriptionFunc),
	}
}

func (eb *InMemEventBus) Publish(events ...events.Event) error {
	for _, event := range events {
		subs, ok := eb.subscriptions[event.Topic().Value()]
		if !ok {
			continue
		}

		for _, sub := range subs {
			if err := sub(event); err != nil {
				return err
			}
		}
	}

	return nil
}

func (eb *InMemEventBus) Subscribe(fn events.SubscriptionFunc, topics ...events.Topic) {
	for _, topic := range topics {
		subs, ok := eb.subscriptions[topic.Value()]
		if !ok {
			subs = make([]events.SubscriptionFunc, 0)
		}

		subs = append(subs, fn)

		eb.subscriptions[topic.Value()] = subs
	}
}
