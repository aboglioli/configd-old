package schema

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aboglioli/configd/common/model"
	"github.com/aboglioli/configd/domain/config"
	"github.com/aboglioli/configd/domain/props"
)

type Schema struct {
	slug  *model.Slug
	name  *Name
	props map[string]props.Prop
}

func BuildSchema(slug *model.Slug, name *Name, ps ...props.Prop) (*Schema, error) {
	if len(ps) == 0 {
		return nil, errors.New("schema does not have props")
	}

	psMap := make(map[string]props.Prop)
	for _, p := range ps {
		psMap[p.Name()] = p
	}

	return &Schema{
		slug:  slug,
		name:  name,
		props: psMap,
	}, nil
}

func NewSchema(name *Name, ps ...props.Prop) (*Schema, error) {
	slug, err := model.NewSlug(name.Value())
	if err != nil {
		return nil, err
	}

	return BuildSchema(slug, name, ps...)
}

func (s *Schema) Slug() *model.Slug {
	return s.slug
}

func (s *Schema) Name() *Name {
	return s.name
}

func (s *Schema) Props() map[string]props.Prop {
	return s.props
}

func (s *Schema) Validate(c config.ConfigData) error {
	for k, v := range c {
		if err := validateConfig(k, v, s.props); err != nil {
			return err
		}
	}

	return nil
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	d := map[string]interface{}{
		"slug":  s.slug.Value(),
		"name":  s.name.Value(),
		"props": s.props,
	}

	return json.Marshal(&d)
}

func validateConfig(key string, value interface{}, propsMap map[string]props.Prop) error {
	prop, ok := propsMap[key]
	if !ok {
		return fmt.Errorf("prop %s not found in schema", key)
	}

	if value == nil && prop.Default() != nil {
		value = prop.Default()
	}

	if value == nil && prop.Required() {
		return fmt.Errorf("value for %s is required", key)
	}

	if len(prop.Enum()) > 0 {
		isFromEnum := false
		for _, e := range prop.Enum() {
			if value == e {
				isFromEnum = true
				break
			}
		}

		if !isFromEnum {
			return fmt.Errorf("%s:%v is not in enum values", key, value)
		}
	}

	switch value := value.(type) {
	case string:
		if prop.Type() != props.STRING {
			return fmt.Errorf("%v is not a string", value)
		}
	case int:
		if prop.Type() != props.INT {
			return fmt.Errorf("%s:%v is not an integer", key, value)
		}

		if prop.Interval() != nil {
			interval := prop.Interval()

			if value < int(interval.Min()) {
				return fmt.Errorf("%s:%v is lesser than the minimum value in interval", key, value)
			}

			if value > int(interval.Max()) {
				return fmt.Errorf("%s:%v is greater than the maximum value in interval", key, value)
			}
		}
	case int32:
		if prop.Type() != props.INT {
			return fmt.Errorf("%s:%v is not an integer", key, value)
		}

		if prop.Interval() != nil {
			interval := prop.Interval()

			if value < int32(interval.Min()) {
				return fmt.Errorf("%s:%v is lesser than the minimum value in interval", key, value)
			}

			if value > int32(interval.Max()) {
				return fmt.Errorf("%s:%v is greater than the maximum value in interval", key, value)
			}
		}
	case int64:
		if prop.Type() != props.INT {
			return fmt.Errorf("%s:%v is not an integer", key, value)
		}

		if prop.Interval() != nil {
			interval := prop.Interval()

			if value < int64(interval.Min()) {
				return fmt.Errorf("%s:%v is lesser than the minimum value in interval", key, value)
			}

			if value > int64(interval.Max()) {
				return fmt.Errorf("%s:%v is greater than the maximum value in interval", key, value)
			}
		}
	case float32:
		if prop.Type() != props.FLOAT {
			return fmt.Errorf("%s:%v is not a float", key, value)
		}

		if prop.Interval() != nil {
			interval := prop.Interval()

			if value < float32(interval.Min()) {
				return fmt.Errorf("%s:%v is lesser than the minimum value in interval", key, value)
			}

			if value > float32(interval.Max()) {
				return fmt.Errorf("%s:%v is greater than the maximum value in interval", key, value)
			}
		}
	case float64:
		if prop.Type() != props.FLOAT {
			return fmt.Errorf("%s:%v is not a float", key, value)
		}

		if prop.Interval() != nil {
			interval := prop.Interval()

			if value < interval.Min() {
				return fmt.Errorf("%s:%v is lesser than the minimum value in interval", key, value)
			}

			if value > interval.Max() {
				return fmt.Errorf("%s:%v is greater than the maximum value in interval", key, value)
			}
		}
	case bool:
		if prop.Type() != props.BOOL {
			return fmt.Errorf("%s:%v is not a boolean", key, value)
		}
	case map[string]interface{}:
		if prop.Type() != props.OBJECT {
			return fmt.Errorf("%s:%v is not an object", key, value)
		}

		for k, v := range value {
			if err := validateConfig(k, v, prop.Props()); err != nil {
				return fmt.Errorf("object %s: %s", key, err.Error())
			}
		}
	case []interface{}:
		if !prop.IsArray() {
			return fmt.Errorf("%s is not an array", key)
		}

		for i, v := range value {
			if err := validateConfig(key, v, propsMap); err != nil {
				return fmt.Errorf("array item %d: %s", i, err.Error())
			}
		}
	default:
		return fmt.Errorf("%v unknown type in %s", value, key)
	}

	return nil
}
