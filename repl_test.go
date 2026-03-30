package main

import (
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
		{
			input:    " HELLO WORLD AGAIN",
			expected: []string{"HELLO", "WORLD", "AGAIN"},
		},
		{
			input:    "HELLO XYZ",
			expected: []string{"HELLO", "XYZ"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input %q: got %d words, want %d", c.input, len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Content mismatch at index %d for input %q: got %q, want %q", i, c.input, word, expectedWord)
			}
		}
	}
}
