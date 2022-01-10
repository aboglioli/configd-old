package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id        string
	aggId     string
	topic     Topic
	payload   interface{}
	timestamp time.Time
	version   uint
}

func BuildEvent(
	id string,
	aggId string,
	topic Topic,
	payload interface{},
	timestamp time.Time,
	version uint,
) (Event, error) {
	return Event{
		id:        id,
		aggId:     aggId,
		topic:     topic,
		payload:   payload,
		timestamp: timestamp,
	}, nil
}

func NewEvent(
	aggId string,
	topic Topic,
	payload interface{},
) (Event, error) {
	return BuildEvent(
		uuid.NewString(),
		aggId,
		topic,
		payload,
		time.Now(),
		1,
	)
}

func (e Event) Id() string {
	return e.id
}

func (e Event) AggregateRootId() string {
	return e.aggId
}

func (e Event) Topic() Topic {
	return e.topic
}

func (e Event) Payload() interface{} {
	return e.payload
}

func (e Event) Timestamp() time.Time {
	return e.timestamp
}

func (e Event) Version() uint {
	return e.version
}
