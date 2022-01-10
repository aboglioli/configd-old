package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type Id struct {
	id string
}

func BuildId(id string) (Id, error) {
	if len(id) < 4 {
		return Id{}, errors.New("id too short")
	}

	return Id{
		id: id,
	}, nil
}

func NewUuid() (Id, error) {
	return BuildId(uuid.NewString())
}

func NewSlug(str string) (Id, error) {
	return BuildId(slug.Make(str))
}

func (id Id) Value() string {
	return id.id
}

func (id Id) Equals(o Id) bool {
	return id.id == o.id
}
