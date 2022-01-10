package events

type EventPublisher interface {
	Publish(events ...Event) error
}
