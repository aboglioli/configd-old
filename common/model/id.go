package model

import (
	"errors"

	"github.com/google/uuid"
)

type Id struct {
	id string
}

func NewId(id string) (*Id, error) {
	if len(id) < 4 {
		return nil, errors.New("id too short")
	}

	return &Id{
		id: id,
	}, nil
}

func GenerateId() (*Id, error) {
	return NewId(uuid.NewString())
}
