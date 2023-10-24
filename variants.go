package variants

import (
	"reflect"
	"strings"
)

type Options[T comparable] struct {
	Variants         map[string]map[any]string
	DefaultVariants  map[string]any
	CompoundVariants map[T]string
}

func New[T comparable](base string, options Options[T]) func(T) string {
	return func(props T) string {
		classNames := classList{base}

		p := reflect.ValueOf(props)
		for variant, mapping := range options.Variants {
			switch p.Kind() {
			case reflect.Map:
				if value, ok := p.MapIndex(reflect.ValueOf(variant)).Interface().(any); ok {
					classNames.Add(mapping[value])
				} else {
					classNames.Add(mapping[options.DefaultVariants[variant]])
				}
			case reflect.Struct:
				if value, ok := p.FieldByName(variant).Interface().(any); ok {
					classNames.Add(mapping[value])
				} else {
					classNames.Add(mapping[options.DefaultVariants[variant]])
				}
			default:
				classNames.Add(mapping[options.DefaultVariants[variant]])
			}
		}

		for match, className := range options.CompoundVariants {
			matchValue := reflect.ValueOf(match)
			if matchValue.Kind() == reflect.Map {
				if matchValue.Len() == 0 {
					continue
				}
				for _, key := range matchValue.MapKeys() {
					if matchValue.MapIndex(key).Interface() != p.MapIndex(key).Interface() {
						continue
					}
				}
			}

			if matchValue.Kind() == reflect.Struct {
				for i := 0; i < matchValue.NumField(); i++ {
					if matchValue.Field(i).Interface() != p.Field(i).Interface() {
						continue
					}
				}
			}

			classNames.Add(className)
		}

		return classNames.String()
	}
}

type classList []string

func (c classList) String() string {
	return strings.TrimSpace(strings.Join(c, " "))
}

func (c *classList) Add(class string) {
	if class == "" {
		return
	}
	*c = append(*c, class)
}
