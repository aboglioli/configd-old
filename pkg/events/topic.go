package events

import (
	"strings"
)

type Topic struct {
	topic string
}

func NewTopic(path ...string) Topic {
	if len(path) == 0 {
		panic("empty topic")
	}

	for _, p := range path {
		if p == "" {
			panic("empty topic path")
		}
	}

	return Topic{
		topic: strings.Join(path, "."),
	}
}

func (t Topic) Value() string {
	return t.topic
}

func (t Topic) Equals(o Topic) bool {
	return t.topic == o.topic
}
