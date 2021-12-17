package schema

import (
	"errors"

	"github.com/aboglioli/configd/domain/props"
)

type Schema []props.Prop

func NewSchema(props ...props.Prop) (Schema, error) {
	if len(props) == 0 {
		return nil, errors.New("schema does not have props")
	}

	return props, nil
}
