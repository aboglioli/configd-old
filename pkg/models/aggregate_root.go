package models

import (
	"fmt"
	"time"

	"github.com/aboglioli/configd/pkg/events"
)

type AggregateRoot struct {
	events    []*events.Event
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func BuildAggregateRoot(
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*AggregateRoot, error) {
	if updatedAt.Sub(createdAt) < 0 {
		return nil, fmt.Errorf(
			"createdAt %s timestamp is greater than updatedAt %s",
			createdAt.Format(time.RFC3339),
			updatedAt.Format(time.RFC3339),
		)
	}

	return &AggregateRoot{
		events:    make([]*events.Event, 0),
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}, nil
}

func NewAggregateRoot() (*AggregateRoot, error) {
	return BuildAggregateRoot(time.Now(), time.Now(), nil)
}

func (a *AggregateRoot) Events() []*events.Event {
	return a.events
}

func (a *AggregateRoot) RecordEvents(events ...*events.Event) {
	a.events = append(a.events, events...)
}

func (a *AggregateRoot) CreatedAt() time.Time {
	return a.createdAt
}

func (a *AggregateRoot) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *AggregateRoot) DeletedAt() *time.Time {
	return a.deletedAt
}
