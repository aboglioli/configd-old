package props

import (
	"errors"
)

type Prop interface {
	Name() string
	Type() PropType
	Default() interface{}
	Required() bool
	Values() []interface{}
	Regex() string
	Interval() *interval
	Props() map[string]Prop
}

var _ Prop = (*prop)(nil)

type prop struct {
	name     string
	t        PropType
	def      interface{}
	required bool
	values   []interface{}
	regex    string
	interval *interval
	props    map[string]Prop
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

func (p *prop) Values() []interface{} {
	return p.values
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
