package models

import (
	"errors"

	"github.com/gosimple/slug"
)

type Slug struct {
	slug string
}

func NewSlug(s string) (*Slug, error) {
	if len(s) == 0 {
		return nil, errors.New("empty slug")
	}

	s = slug.Make(s)

	return &Slug{
		slug: s,
	}, nil
}

func (s *Slug) Value() string {
	return s.slug
}

func (s *Slug) Equals(o *Slug) bool {
	return s.slug == o.slug
}
