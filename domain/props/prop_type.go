package props

import (
	"errors"
)

type PropType string

const (
	STRING PropType = "string"
	INT             = "integer"
	FLOAT           = "float"
	BOOL            = "bool"
	OBJECT          = "object"
)

func (t PropType) IsValid() bool {
	if t == STRING || t == INT || t == FLOAT || t == BOOL || t == OBJECT {
		return true
	}
	return false
}

func NewPropType(t string) (PropType, error) {
	switch t {
	case string(STRING):
		return STRING, nil
	case string(INT):
		return INT, nil
	case string(FLOAT):
		return FLOAT, nil
	case string(BOOL):
		return BOOL, nil
	case string(OBJECT):
		return OBJECT, nil
	default:
		return "", errors.New("invalid type")
	}
}
