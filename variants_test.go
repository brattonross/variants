package variants

import (
	"slices"
	"strings"
	"testing"
)

func TestVariants(t *testing.T) {
	tt := []struct {
		name     string
		base     string
		options  Options[any]
		props    any
		expected string
	}{
		{
			name: "no defaults",
			base: "base",
			options: Options[any]{
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
			},
			props:    nil,
			expected: "base",
		},
		{
			name: "with defaults",
			base: "base",
			options: Options[any]{
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
			},
			props:    nil,
			expected: "base button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600 button--enabled cursor-pointer",
		},
		{
			name: "with defaults and overrides",
			base: "base",
			options: Options[any]{
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
			},
			props: map[string]any{
				"Intent":   "secondary",
				"Disabled": true,
			},
			expected: "base button--secondary bg-white text-gray-800 border-gray-400 hover:bg-gray-100 button--disabled opacity-50 cursor-not-allowed",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, newAssertionFunc(tc.base, tc.options, tc.props, tc.expected))
	}

	t.Run("compound variants", func(t *testing.T) {
		type Props struct {
			Intent   string
			Disabled bool
		}
		newAssertionFunc(
			"base",
			Options[Props]{
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
			},
			Props{
				Intent:   "primary",
				Disabled: false,
			},
			"base button--primary bg-blue-500 text-white border-transparent hover:bg-blue-600 button--enabled cursor-pointer button--primary-medium uppercase",
		)(t)
	})
}

func newAssertionFunc[T comparable](base string, options Options[T], props T, expected string) func(*testing.T) {
	return func(t *testing.T) {
		fn := New(base, options)
		actual := fn(props)

		ap := strings.Split(actual, " ")
		ep := strings.Split(expected, " ")

		slices.Sort(ap)
		slices.Sort(ep)

		for i := range ap {
			if ap[i] != ep[i] {
				t.Errorf("got %s, want %s", actual, expected)
			}
		}
	}
}

func TestCx(t *testing.T) {
	tt := []struct {
		name     string
		classes  []any
		expected string
	}{
		{
			name:     "empty",
			classes:  []any{},
			expected: "",
		},
		{
			name:     "single",
			classes:  []any{"foo"},
			expected: "foo",
		},
		{
			name:     "multiple",
			classes:  []any{"foo", "bar"},
			expected: "foo bar",
		},
		{
			name:     "multiple with empty",
			classes:  []any{"foo", "", "bar"},
			expected: "foo bar",
		},
		{
			name: "nested",
			classes: []any{
				"foo",
				[]string{"bar", "baz"},
				[]any{"qux", []string{"quux", "corge"}},
			},
			expected: "foo bar baz qux quux corge",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := Cx(tc.classes...)
			if actual != tc.expected {
				t.Errorf("got %s, want %s", actual, tc.expected)
			}
		})
	}
}
