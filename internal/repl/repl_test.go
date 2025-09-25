package repl

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 		string
		expected	[]string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " hello mother ",
			expected: []string{"hello", "mother"},
		},
		{
			input: "xuxubinha  is Xuxubistica",
			expected: []string{"xuxubinha", "is", "xuxubistica"},
		},
		{
			input: "    BOBOZILDO ReI da BObolandia",
			expected: []string{"bobozildo", "rei", "da", "bobolandia"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Expected length %d, but got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual: %s but expected: %s", word, expectedWord)
			}
		}
	}
	
}