package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello    world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  tEa  coFfee   water",
			expected: []string{"tea", "coffee", "water"},
		},
		{
			input:    "\n\n hello frIEND\n\n",
			expected: []string{"hello", "friend"},
		},
		{
			input:    " i like \ncoIl anD \nautechre",
			expected: []string{"i", "like", "coil", "and", "autechre"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v; got: %v", c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if !reflect.DeepEqual(word, expectedWord) {
				t.Errorf("expected string: %v; got: %v", expectedWord, word)
			}
		}
	}
}
