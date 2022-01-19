package user

import (
	"fmt"
)

type Access string

const (
	READ_ONLY_ACCESS Access = "read_only"
	FULL_ACCESS      Access = "full_access"
)

func NewAccess(access string) (Access, error) {
	switch access {
	case string(READ_ONLY_ACCESS):
		return READ_ONLY_ACCESS, nil
	case string(FULL_ACCESS):
		return FULL_ACCESS, nil
	}

	return "", fmt.Errorf("invalid access %s", access)
}
