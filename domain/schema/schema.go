package schema

import (
	"errors"

	"github.com/aboglioli/configd/domain/props"
)

type Schema map[string]props.Prop

func NewSchema(props ...props.Prop) (Schema, error) {
	if len(props) == 0 {
		return nil, errors.New("schema does not have props")
	}

	s := make(Schema)
	for _, p := range props {
		s[p.Name()] = p
	}

	return s, nil
}
