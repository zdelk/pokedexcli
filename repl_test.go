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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "	This 	iS	A	tEsT",
			expected: []string{"this", "is", "a", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(`------------------------------
			Input: 		(%v)
			Expecting: 	%s
			Actual:		%s
			Fail 
			`, c.input, c.expected, actual)
		}

		for i := range actual {

			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(`--------------------------
				Expecting:	%s
				Actual:		%s
				Fail
				`, expectedWord, word)
			}
		}
	}
}
