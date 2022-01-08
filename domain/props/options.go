package props

import (
	"errors"
	"fmt"
)

type Option func(p *Prop) error

func WithDefault(d interface{}) Option {
	return func(p *Prop) error {
		if p.t == OBJECT {
			return fmt.Errorf("%s cannot have default", p.t)
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
	return func(p *Prop) error {
		if p.t == OBJECT && !r {
			return fmt.Errorf("%s must be required", p.t)
		}

		p.required = r
		return nil
	}
}

func WithEnum(enum ...interface{}) Option {
	return func(p *Prop) error {
		if p.t == OBJECT {
			return fmt.Errorf("%s cannot have enum values", p.t)
		}

		parsedEnum := make([]interface{}, len(enum))
		for i, value := range enum {
			switch p.t {
			case STRING:
				if _, ok := value.(string); !ok {
					return errors.New("string enum expected")
				}
			case INT:
				// Convert between number types
				if v, ok := value.(float64); ok {
					value = int(v)
				}

				if _, ok := value.(int); !ok {
					return errors.New("integer enum expected")
				}
			case FLOAT:
				// Convert between number types
				if v, ok := value.(int); ok {
					value = float64(v)
				}

				if _, ok := value.(float64); !ok {
					return errors.New("float enum expected")
				}
			case BOOL:
				if _, ok := value.(bool); !ok {
					return errors.New("bool enum expected")
				}
			}

			parsedEnum[i] = value
		}

		p.enum = parsedEnum
		return nil
	}
}

func WithRegex(regex string) Option {
	return func(p *Prop) error {
		if p.t != STRING {
			return fmt.Errorf("%s cannot have regex", p.t)
		}

		p.regex = regex
		return nil
	}
}

func WithInterval(min, max float64) Option {
	return func(p *Prop) error {
		if p.t != INT && p.t != FLOAT {
			return fmt.Errorf("%s cannot have interval, it must be used with numeric types", p.t)
		}

		interval, err := NewInterval(min, max)
		if err != nil {
			return err
		}

		p.interval = interval
		return nil
	}
}

func WithProps(props ...*Prop) Option {
	return func(p *Prop) error {
		if p.t != OBJECT {
			return fmt.Errorf("%s cannot have subprops", p.t)
		}

		if p.props == nil {
			p.props = make(map[string]*Prop)
		}

		for _, prop := range props {
			p.props[prop.Name()] = prop
		}

		return nil
	}
}

func WithArray(props ...*Prop) Option {
	return func(p *Prop) error {
		p.array = true
		return nil
	}
}
