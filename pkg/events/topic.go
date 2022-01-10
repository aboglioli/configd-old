package events

import (
	"errors"
	"strings"
)

type topic struct {
	topic string
}

func NewTopic(path ...string) (*topic, error) {
	if len(path) == 0 {
		return nil, errors.New("empty topic")
	}

	for _, p := range path {
		if p == "" {
			return nil, errors.New("empty topic path")
		}
	}

	return &topic{
		topic: strings.Join(path, "."),
	}, nil
}

func (t *topic) Value() string {
	return t.topic
}

func (t *topic) Equals(o *topic) bool {
	return t.topic == o.topic
}
