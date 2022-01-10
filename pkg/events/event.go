package events

import (
	"time"

	"github.com/google/uuid"
)

type event struct {
	id        string
	aggId     string
	topic     *topic
	payload   interface{}
	timestamp time.Time
	version   uint
}

func BuildEvent(
	id string,
	aggId string,
	topic *topic,
	payload interface{},
	timestamp time.Time,
	version uint,
) (*event, error) {
	return &event{
		id:        id,
		aggId:     aggId,
		topic:     topic,
		payload:   payload,
		timestamp: timestamp,
	}, nil
}

func NewEvent(
	aggId string,
	topic *topic,
	payload interface{},
) (*event, error) {
	return BuildEvent(
		uuid.NewString(),
		aggId,
		topic,
		payload,
		time.Now(),
		1,
	)
}

func (e *event) Id() string {
	return e.id
}

func (e *event) AggregateRootId() string {
	return e.aggId
}

func (e *event) Topic() string {
	return e.topic.Value()
}

func (e *event) Payload() interface{} {
	return e.payload
}

func (e *event) Timestamp() time.Time {
	return e.timestamp
}

func (e *event) Version() uint {
	return e.version
}

func (e *event) WithVersion(version uint) *event {
	e.version = version
	return e
}
