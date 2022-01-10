package models

import (
	"time"
)

type Event interface {
	Id() string
	AggregateRootId() string
	Topic() string
	Payload() interface{}
	Timestamp() time.Time
	Version() uint
}
