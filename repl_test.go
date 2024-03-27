package main

import "testing"

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, test := range testCases {
		actual := cleanInput(test.input)

		if len(actual) != len(test.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(test.expected), len(actual))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := test.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}

}
