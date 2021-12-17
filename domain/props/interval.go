package props

import (
	"errors"
)

type interval struct {
	min float64
	max float64
}

func NewInterval(min, max float64) (*interval, error) {
	if min > max {
		return nil, errors.New("invalid interval")
	}

	return &interval{min, max}, nil
}

func (i *interval) Min() float64 {
	return i.min
}

func (i *interval) Max() float64 {
	return i.max
}
