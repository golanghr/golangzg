package main_test

import (
	"testing"

	main "github.com/golanghr/golangzg/hacktable/19-01/matejb/subtests"
)

func TestTheThing(t *testing.T) {
	tt := []struct {
		it       string
		name     string
		expected string
	}{
		{
			it:       "returns the thing with name",
			name:     "Go",
			expected: "the thing Go",
		},

		{
			it:       "fails miserably",
			name:     "Go",
			expected: "the thing WUT",
		},
	}

	for _, c := range tt {
		t.Run(c.it, func(t *testing.T) {
			result := main.TheThing(c.name)
			if result != c.expected {
				t.Errorf("expected %q got %q", c.expected, result)
			}
		})
	}
}
