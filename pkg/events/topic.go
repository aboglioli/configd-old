package events

import (
	"errors"
	"strings"
)

type Topic struct {
	topic string
}

func NewTopic(path ...string) (Topic, error) {
	if len(path) == 0 {
		return Topic{}, errors.New("empty topic")
	}

	for _, p := range path {
		if p == "" {
			return Topic{}, errors.New("empty topic path")
		}
	}

	return Topic{
		topic: strings.Join(path, "."),
	}, nil
}

func (t Topic) Value() string {
	return t.topic
}

func (t Topic) Equals(o Topic) bool {
	return t.topic == o.topic
}
