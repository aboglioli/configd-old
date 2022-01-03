package props

import (
	"encoding/json"
	"errors"
)

type Prop interface {
	Name() string
	Type() PropType
	Default() interface{}
	Required() bool
	Enum() []interface{}
	Regex() string
	Interval() *interval
	Props() map[string]Prop
	IsArray() bool
}

var _ Prop = (*prop)(nil)

type prop struct {
	name     string
	t        PropType
	def      interface{}
	required bool
	enum     []interface{}
	regex    string
	interval *interval
	props    map[string]Prop
	array    bool
}

func newValue(name string, t PropType, opts ...Option) (*prop, error) {
	if name == "" {
		return nil, errors.New("empty name in prop")
	}

	p := &prop{
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

func NewString(name string, opts ...Option) (*prop, error) {
	return newValue(name, STRING, opts...)
}

func NewInteger(name string, opts ...Option) (*prop, error) {
	return newValue(name, INT, opts...)
}

func NewFloat(name string, opts ...Option) (*prop, error) {
	return newValue(name, FLOAT, opts...)
}

func NewBool(name string, opts ...Option) (*prop, error) {
	return newValue(name, BOOL, opts...)
}

func NewObject(name string, opts ...Option) (*prop, error) {
	// Objects should be required always
	opts = append(opts, WithRequired(true))

	return newValue(
		name,
		OBJECT,
		opts...,
	)
}

func (p *prop) Name() string {
	return p.name
}

func (p *prop) Type() PropType {
	return p.t
}

func (p *prop) Default() interface{} {
	return p.def
}

func (p *prop) Required() bool {
	return p.required
}

func (p *prop) Enum() []interface{} {
	return p.enum
}

func (p *prop) Regex() string {
	return p.regex
}

func (p *prop) Interval() *interval {
	return p.interval
}

func (p *prop) Props() map[string]Prop {
	return p.props
}

func (p *prop) IsArray() bool {
	return p.array
}

func (p *prop) MarshalJSON() ([]byte, error) {
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
