package main

import(
	"testing"
)

func TestCleanInput(t* testing.t) {
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
        input:    "   ",
        expected: []string{},
    },
    {
        input:    "one      two   three",
        expected: []string{"one", "two", "three"},
    },
    // You can add more cases if you wish
}
}
for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
		continue
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}