package props

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Prop struct {
	name     string
	t        PropType
	def      interface{}
	required bool
	enum     []interface{}
	regex    string
	interval *Interval
	props    map[string]*Prop
	array    bool
}

func newValue(name string, t PropType, opts ...Option) (*Prop, error) {
	if name == "" {
		return nil, errors.New("empty name in prop")
	}

	p := &Prop{
		t:    t,
		name: name,
	}

	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}

	return p, nil
}

func NewString(name string, opts ...Option) (*Prop, error) {
	return newValue(name, STRING, opts...)
}

func NewInteger(name string, opts ...Option) (*Prop, error) {
	return newValue(name, INT, opts...)
}

func NewFloat(name string, opts ...Option) (*Prop, error) {
	return newValue(name, FLOAT, opts...)
}

func NewBool(name string, opts ...Option) (*Prop, error) {
	return newValue(name, BOOL, opts...)
}

func NewObject(name string, opts ...Option) (*Prop, error) {
	// Objects should be required always
	opts = append(opts, WithRequired())

	return newValue(
		name,
		OBJECT,
		opts...,
	)
}

func (p *Prop) Name() string {
	return p.name
}

func (p *Prop) Type() PropType {
	return p.t
}

func (p *Prop) Default() interface{} {
	return p.def
}

func (p *Prop) IsRequired() bool {
	return p.required
}

func (p *Prop) Enum() []interface{} {
	return p.enum
}

func (p *Prop) Regex() string {
	return p.regex
}

func (p *Prop) Interval() *Interval {
	return p.interval
}

func (p *Prop) Props() map[string]*Prop {
	return p.props
}

func (p *Prop) IsArray() bool {
	return p.array
}

func (p *Prop) MarshalJSON() ([]byte, error) {
	d := map[string]interface{}{
		"name":     p.name,
		"type":     p.t,
		"default":  p.def,
		"required": p.required,
		"enum":     p.enum,
		"regex":    p.regex,
		"props":    p.props,
		"array":    p.array,
	}

	if p.interval != nil {
		d["interval"] = map[string]interface{}{
			"min": p.interval.min,
			"max": p.interval.max,
		}
	}

	return json.Marshal(&d)
}

func (p *Prop) Validate(v interface{}) error {
	return p.validateWithArray(v, true)
}

func (p *Prop) validateWithArray(v interface{}, validateArray bool) error {
	// Required
	if p.IsRequired() && v == nil {
		return errors.New("value is required")
	}

	// Validate array elements
	if validateArray && p.IsArray() {
		arr, ok := v.([]interface{})
		if !ok {
			return fmt.Errorf("%v is not an array", v)
		}

		for i, v := range arr {
			if err := p.validateWithArray(v, false); err != nil {
				return fmt.Errorf("array item %d: %s", i, err.Error())
			}
		}

		return nil
	}

	// Check for enum values
	if len(p.Enum()) > 0 {
		isInEnum := false
		for _, e := range p.Enum() {
			if v == e {
				isInEnum = true
				break
			}
		}

		if !isInEnum {
			return fmt.Errorf("%v is not in enum values %v", v, p.Enum())
		}
	}

	switch p.Type() {
	case STRING:
		_, ok := v.(string)
		if !ok {
			return fmt.Errorf("%v is not a string", v)
		}
	case INT:
		i, okInt := v.(int)
		i32, okInt32 := v.(int32)
		i64, okInt64 := v.(int64)

		f32, okFloat32 := v.(float32)
		f64, okFloat64 := v.(float64)

		if !okInt {
			if okInt32 {
				i = int(i32)
			} else if okInt64 {
				i = int(i64)
			} else if okFloat32 {
				i = int(f32)
			} else if okFloat64 {
				i = int(f64)
			}else {
				return fmt.Errorf("%v is not an integer", v)
			}
		}

		if p.Interval() != nil {
			interval := p.Interval()

			if i < int(interval.Min()) {
				return fmt.Errorf("%v is lesser than the minimum value in interval", v)
			}

			if i > int(interval.Max()) {
				return fmt.Errorf("%v is greater than the maximum value in interval", v)
			}
		}
	case FLOAT:
		f32, okFloat32 := v.(float32)
		f, okFloat64 := v.(float64)

		if !okFloat64 {
			if okFloat32 {
				f = float64(f32)
			} else {
				return fmt.Errorf("%v is not a float", v)
			}
		}

		if p.Interval() != nil {
			interval := p.Interval()

			if f < float64(interval.Min()) {
				return fmt.Errorf("%v is lesser than the minimum value in interval", v)
			}

			if f > float64(interval.Max()) {
				return fmt.Errorf("%v is greater than the maximum value in interval", v)
			}
		}
	case BOOL:
		_, ok := v.(bool)
		if !ok {
			return fmt.Errorf("%v is not a boolean", v)
		}
	case OBJECT:
		obj, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%v is not an object", v)
		}

		if len(p.Props()) == 0 {
			return fmt.Errorf("%v does not have subprops", v)
		}

		for k, p := range p.Props() {
			v, ok := obj[k]
			if !ok {
				return fmt.Errorf("missing prop for key %s", k)
			}

			if err := p.validateWithArray(v, true); err != nil {
				return fmt.Errorf("key %s: %s", k, err.Error())
			}
		}
	}

	return nil
}
