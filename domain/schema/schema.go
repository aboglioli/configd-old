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
	// for k, v := range c {
	// 	if err := walkConfig(k, v, s.props); err != nil {
	// 		return err
	// 	}
	// }

	// return nil

	return walkProps(s.props, c, true)
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	d := map[string]interface{}{
		"slug":  s.slug.Value(),
		"name":  s.name.Value(),
		"props": s.props,
	}

	return json.Marshal(&d)
}

func walkConfig(key string, value interface{}, propsMap map[string]props.Prop) error {
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
			if err := walkConfig(k, v, prop.Props()); err != nil {
				return fmt.Errorf("object %s: %s", key, err.Error())
			}
		}
	case []interface{}:
		if !prop.IsArray() {
			return fmt.Errorf("%s is not an array", key)
		}

		for i, v := range value {
			if err := walkConfig(key, v, propsMap); err != nil {
				return fmt.Errorf("array item %d: %s", i, err.Error())
			}
		}
	default:
		return fmt.Errorf("%v unknown type in %s", value, key)
	}

	return nil
}

func walkProps(propsMap map[string]props.Prop, config map[string]interface{}, validateArray bool) error {
	for k, prop := range propsMap {
		entry, ok := config[k]
		if !ok {
			return fmt.Errorf("prop %s not found in config", k)
		}

		// Required and assign default value
		if prop.Required() && entry == nil {
			if prop.Default() != nil {
				config[k] = prop.Default()
			} else {
				return fmt.Errorf("value for %s is required", k)
			}
		}

		// Check for enum values
		if len(prop.Enum()) > 0 {
			isInEnum := false
			for _, e := range prop.Enum() {
				if entry == e {
					isInEnum = true
					break
				}
			}

			if !isInEnum {
				return fmt.Errorf("%v from %s is not in enum values %v", entry, k, prop.Enum())
			}
		}

		if prop.IsArray() && validateArray {
			arr, ok := entry.([]interface{})
			if !ok {
				return fmt.Errorf("%s is not an array", k)
			}

			for i, v := range arr {
				if err := walkProps(propsMap, map[string]interface{}{k: v}, false); err != nil {
					return fmt.Errorf("array item %d: %s", i, err.Error())
				}
			}

			return nil
		}

		switch prop.Type() {
		case props.STRING:
			_, ok := entry.(string)
			if !ok {
				return fmt.Errorf("%v is not a string", entry)
			}
		case props.INT:
			v, okInt := entry.(int)
			v32, okInt32 := entry.(int32)
			v64, okInt64 := entry.(int64)

			if !okInt {
				if okInt32 {
					v = int(v32)
				} else if okInt64 {
					v = int(v64)
				} else {
					return fmt.Errorf("%s = %v is not an integer", k, entry)
				}
			}

			if prop.Interval() != nil {
				interval := prop.Interval()

				if v < int(interval.Min()) {
					return fmt.Errorf("%s = %v is lesser than the minimum value in interval", k, entry)
				}

				if v > int(interval.Max()) {
					return fmt.Errorf("%s = %v is greater than the maximum value in interval", k, entry)
				}
			}
		case props.FLOAT:
			v32, okFloat32 := entry.(float32)
			v, okFloat64 := entry.(float64)

			if !okFloat64 {
				if okFloat32 {
					v = float64(v32)
				} else {
					return fmt.Errorf("%s = %v is not a float", k, entry)
				}
			}

			if prop.Interval() != nil {
				interval := prop.Interval()

				if v < float64(interval.Min()) {
					return fmt.Errorf("%s = %v is lesser than the minimum value in interval", k, entry)
				}

				if v > float64(interval.Max()) {
					return fmt.Errorf("%s = %v is greater than the maximum value in interval", k, entry)
				}
			}
		case props.BOOL:
			_, ok := entry.(bool)
			if !ok {
				return fmt.Errorf("%s = %v is not a boolean", k, entry)
			}
		case props.OBJECT:
			obj, ok := entry.(map[string]interface{})
			if !ok {
				return fmt.Errorf("%s = %v is not an object", k, entry)
			}

			if err := walkProps(prop.Props(), obj, true); err != nil {
				return fmt.Errorf("object %s: %s", k, err.Error())
			}
		}
	}

	return nil
}
