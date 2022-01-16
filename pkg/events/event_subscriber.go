package events

type SubscriptionFunc func(evt Event) error

type EventSubscriber interface {
	Subscribe(fn SubscriptionFunc, topics ...Topic)
}
