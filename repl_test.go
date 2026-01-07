package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	splitTests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Case 1: two words with whitespaces",
			input:    "  hello       world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Case 2: empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Case 3: one word with whitespaces",
			input:    "	hello   ",
			expected: []string{"hello"},
		},
		{
			name:     "Case 4: should lowercase words",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, st := range splitTests {
		t.Run(st.name, func(t *testing.T) {
			actual := cleanInput(st.input)

			got := len(actual)
			want := len(st.expected)

			if got != want {
				t.Errorf("got %v words, want %v", len(actual), len(st.expected))
			}

			for i := range actual {
				got := actual[i]
				want := st.expected[i]
				if got != want {
					t.Errorf("actual %s expected %s", got, want)
				}
			}
		})
	}
}
