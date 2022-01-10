package props

import (
	"errors"
)

type PropType string

const (
	STRING PropType = "string"
	INT    PropType = "integer"
	FLOAT  PropType = "float"
	BOOL   PropType = "bool"
	OBJECT PropType = "object"
)

func (t PropType) IsValid() bool {
	if t == STRING || t == INT || t == FLOAT || t == BOOL || t == OBJECT {
		return true
	}
	return false
}

func (t PropType) String() string {
	return string(t)
}

func NewPropType(t string) (PropType, error) {
	switch t {
	case STRING.String():
		return STRING, nil
	case INT.String():
		return INT, nil
	case FLOAT.String():
		return FLOAT, nil
	case BOOL.String():
		return BOOL, nil
	case OBJECT.String():
		return OBJECT, nil
	default:
		return "", errors.New("invalid type")
	}
}
