package props

import (
	"errors"
)

type Interval struct {
	min float64
	max float64
}

func NewInterval(min, max float64) (Interval, error) {
	if min > max {
		return Interval{}, errors.New("invalid interval")
	}

	return Interval{min, max}, nil
}

func (i Interval) Min() float64 {
	return i.min
}

func (i Interval) Max() float64 {
	return i.max
}
