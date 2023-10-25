package variants

import (
	"fmt"
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
		if any(props) == nil {
			return base
		}

		classNames := []string{base}

		pv := reflect.ValueOf(props)
		kind := pv.Kind()
		for variant, mapping := range options.Variants {
			variantClass := mapping[options.DefaultVariants[variant]]
			switch kind {
			case reflect.Map:
				vv := reflect.ValueOf(variant)
				if vv.IsValid() {
					mv := pv.MapIndex(vv)
					if mv.IsValid() {
						value, ok := mv.Interface().(any)
						if ok {
							variantClass = mapping[value]
						}
					}
				}
			case reflect.Struct:
				f := pv.FieldByName(variant)
				if f.IsValid() {
					value, ok := f.Interface().(any)
					if ok {
						variantClass = mapping[value]
					}
				}
			}
			classNames = append(classNames, variantClass)
		}

		for match, className := range options.CompoundVariants {
			if any(match) == nil {
				continue
			}

			mv := reflect.ValueOf(match)
			if mv.IsValid() {
				switch mv.Kind() {
				case reflect.Map:
					for _, key := range mv.MapKeys() {
						if mv.MapIndex(key).Interface() != pv.MapIndex(key).Interface() {
							continue
						}
					}
				case reflect.Struct:
					for i := 0; i < mv.NumField(); i++ {
						if mv.Field(i).Interface() != pv.Field(i).Interface() {
							continue
						}
					}
				}
			}
			classNames = append(classNames, className)
		}

		return Cx(classNames...)
	}
}

func Cx[T any](classNames ...T) string {
	values := []string{}
	add := func(value string) {
		if value != "" {
			values = append(values, value)
		}
	}

	for _, c := range classNames {
		switch v := any(c).(type) {
		case string:
			add(v)
		case []string:
			for _, s := range v {
				add(s)
			}
		case []any:
			for _, a := range v {
				add(Cx(a))
			}
		case fmt.Stringer:
			add(v.String())
		default:
			add(fmt.Sprintf("%v", v))
		}
	}
	return strings.TrimSpace(strings.Join(values, " "))
}
