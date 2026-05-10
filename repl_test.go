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
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "Zubat ZUbAt ZuBaT",
			expected: []string{"zubat", "zubat", "zubat"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		fmt.Println((actual))
		if len(actual) != len(c.expected) {
			t.Errorf("ERROR: expected and result lengths do not match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("ERROR: expected and actual word mismatch")
			}
		}
	}
}
