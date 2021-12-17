package props

import (
	"errors"
	"fmt"
)

type Option func(p *prop) error

func WithDefault(d interface{}) Option {
	return func(p *prop) error {
		if p.t == OBJECT {
			return errors.New(fmt.Sprintf("%s cannot have default", p.t))
		}

		switch p.t {
		case STRING:
			if _, ok := d.(string); !ok {
				return errors.New("string default value expected")
			}
		case INT:
			// Json unmarshal numbers as float64 by default
			if f, ok := d.(float64); ok {
				d = int(f)
			}
			if _, ok := d.(int); !ok {
				return errors.New("integer default value expected")
			}
		case FLOAT:
			if _, ok := d.(float64); !ok {
				return errors.New("float default value expected")
			}
		case BOOL:
			if _, ok := d.(bool); !ok {
				return errors.New("bool default value expected")
			}
		}

		p.def = d
		return nil
	}
}

func WithRequired(r bool) Option {
	return func(p *prop) error {
		if p.t == OBJECT && !r {
			return errors.New(fmt.Sprintf("%s must be required", p.t))
		}

		p.required = r
		return nil
	}
}

func WithValues(values ...interface{}) Option {
	return func(p *prop) error {
		if p.t == OBJECT {
			return errors.New(fmt.Sprintf("%s cannot have enum values", p.t))
		}

		for _, value := range values {
			switch p.t {
			case STRING:
				if _, ok := value.(string); !ok {
					return errors.New("string values expected")
				}
			case INT:
				if _, ok := value.(int); !ok {
					return errors.New("integer values expected")
				}
			case FLOAT:
				if _, ok := value.(float64); !ok {
					return errors.New("float values expected")
				}
			case BOOL:
				if _, ok := value.(bool); !ok {
					return errors.New("bool values expected")
				}
			}
		}

		p.values = values
		return nil
	}
}

func WithRegex(regex string) Option {
	return func(p *prop) error {
		if p.t != STRING {
			return errors.New(fmt.Sprintf("%s cannot have regex", p.t))
		}

		p.regex = regex
		return nil
	}
}

func WithInterval(min, max float64) Option {
	return func(p *prop) error {
		if p.t != INT && p.t != FLOAT {
			return errors.New(fmt.Sprintf("%s cannot have interval, it must be used with numeric types", p.t))
		}

		interval, err := NewInterval(min, max)
		if err != nil {
			return err
		}

		p.interval = interval
		return nil
	}
}

func WithProps(props ...Prop) Option {
	return func(p *prop) error {
		if p.t != OBJECT {
			return errors.New(fmt.Sprintf("%s cannot have subproperties", p.t))
		}

		if p.props == nil {
			p.props = make(map[string]Prop)
		}

		for _, prop := range props {
			p.props[prop.Name()] = prop
		}

		return nil
	}
}
