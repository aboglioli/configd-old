package models

import (
	"errors"

	"github.com/google/uuid"
)

type Id struct {
	id string
}

func NewId(id string) (Id, error) {
	if len(id) < 4 {
		return Id{}, errors.New("id too short")
	}

	return Id{
		id: id,
	}, nil
}

func GenerateId() (Id, error) {
	return NewId(uuid.NewString())
}

func (id Id) Value() string {
	return id.id
}

func (id Id) Equals(o Id) bool {
	return id.id == o.id
}
