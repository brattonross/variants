package variants

import (
	"testing"
)

func TestVariants(t *testing.T) {
	t.Run("no defaults", func(t *testing.T) {
		fn := New("base", Options[any]{
			Variants: map[string]map[any]string{
				"Intent": {
					"primary":   "button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600",
					"secondary": "button--secondary bg-white text-gray-800 border-gray-400 hover:bg-gray-100",
				},
				"Disabled": {
					true:  "button--disabled opacity-50 cursor-not-allowed",
					false: "button--enabled cursor-pointer",
				},
			},
		})

		actual := fn(nil)
		expected := "base"
		if actual != expected {
			t.Errorf("got %s, want %s", actual, expected)
		}
	})

	t.Run("with defaults", func(t *testing.T) {
		fn := New("base", Options[any]{
			Variants: map[string]map[any]string{
				"Intent": {
					"primary":   "button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600",
					"secondary": "button--secondary bg-white text-gray-800 border-gray-400 hover:bg-gray-100",
				},
				"Disabled": {
					true:  "button--disabled opacity-50 cursor-not-allowed",
					false: "button--enabled cursor-pointer",
				},
			},
			DefaultVariants: map[string]any{
				"Intent":   "primary",
				"Disabled": false,
			},
		})

		actual := fn(nil)
		expected := "base button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600 button--enabled cursor-pointer"
		if actual != expected {
			t.Errorf("got %s, want %s", actual, expected)
		}
	})

	t.Run("with defaults and overrides", func(t *testing.T) {
		fn := New("base", Options[any]{
			Variants: map[string]map[any]string{
				"Intent": {
					"primary":   "button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600",
					"secondary": "button--secondary bg-white text-gray-800 border-gray-400 hover:bg-gray-100",
				},
				"Disabled": {
					true:  "button--disabled opacity-50 cursor-not-allowed",
					false: "button--enabled cursor-pointer",
				},
			},
			DefaultVariants: map[string]any{
				"Intent":   "primary",
				"Disabled": false,
			},
		})

		actual := fn(map[string]any{
			"Intent":   "secondary",
			"Disabled": true,
		})
		expected := "base button--secondary bg-white text-gray-800 border-gray-400 hover:bg-gray-100 button--disabled opacity-50 cursor-not-allowed"
		if actual != expected {
			t.Errorf("got %s, want %s", actual, expected)
		}
	})

	t.Run("compound variants", func(t *testing.T) {
		type Props struct {
			Intent   string
			Disabled bool
		}
		fn := New("base", Options[Props]{
			Variants: map[string]map[any]string{
				"Intent": {
					"primary":   "button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600",
					"secondary": "button--secondary bg-white text-gray-800 border-gray-400 hover:bg-gray-100",
				},
				"Disabled": {
					true:  "button--disabled opacity-50 cursor-not-allowed",
					false: "button--enabled cursor-pointer",
				},
			},
			CompoundVariants: map[Props]string{
				{
					Intent:   "primary",
					Disabled: false,
				}: "button--primary-medium uppercase",
			},
		})

		actual := fn(Props{
			Intent:   "primary",
			Disabled: false,
		})
		expected := "base button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600 button--enabled cursor-pointer button--primary-medium uppercase"
		if actual != expected {
			t.Errorf("got %s, want %s", actual, expected)
		}
	})
}
