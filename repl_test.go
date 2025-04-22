package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length: %v, Actual length: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected: %s, Actual: %s", expectedWord, word)
			}
		}
	}
}
