package models

import (
	"fmt"
	"time"

	"github.com/aboglioli/configd/pkg/events"
)

type PublicAggregateRoot interface {
	Id() Id
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
	Events() []events.Event
	Version() uint
}

var _ PublicAggregateRoot = (*AggregateRoot)(nil)

type AggregateRoot struct {
	id Id

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time

	events []events.Event

	version uint
}

func BuildAggregateRoot(
	id Id,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
	version uint,
) (*AggregateRoot, error) {
	if updatedAt.Sub(createdAt) < 0 {
		return nil, fmt.Errorf(
			"createdAt %s timestamp is greater than updatedAt %s",
			createdAt.Format(time.RFC3339),
			updatedAt.Format(time.RFC3339),
		)
	}

	return &AggregateRoot{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
		events:    make([]events.Event, 0),
		version:   version,
	}, nil
}

func NewAggregateRoot(id Id) (*AggregateRoot, error) {
	return BuildAggregateRoot(id, time.Now(), time.Now(), nil, 1)
}

func (a *AggregateRoot) Id() Id {
	return a.id
}

func (a *AggregateRoot) RecordEvent(events ...events.Event) {
	a.events = append(a.events, events...)
}

func (a *AggregateRoot) CreatedAt() time.Time {
	return a.createdAt
}

func (a *AggregateRoot) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *AggregateRoot) Update() {
	a.updatedAt = time.Now()
	a.version++
}

func (a *AggregateRoot) DeletedAt() *time.Time {
	return a.deletedAt
}

func (a *AggregateRoot) Events() []events.Event {
	return a.events
}

func (a *AggregateRoot) Version() uint {
	return a.version
}
