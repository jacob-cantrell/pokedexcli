package main

import "testing"

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
			input:    "   texT teXt2 TEXT1  ",
			expected: []string{"text", "text2", "text1"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("TEST FAILED WITH WORDS: %s and %s", word, expectedWord)
			}
		}
	}
}
