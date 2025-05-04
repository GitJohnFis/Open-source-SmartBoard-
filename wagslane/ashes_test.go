package main

import "testing"

func TestCleanInputAshes(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        //AI generated test cases
        {input: "  hello  world  ", expected: []string{"hello", "world"}},
        {input: "Charmander Bulbasaur PIKACHU", expected: []string{"charmander", "bulbasaur", "pikachu"}},
        {input: "   ", expected: []string{}},
        {input: "one      two   three", expected: []string{"one", "two", "three"}},
        {input: "hello hello world", expected: []string{"hello", "world"}}, // Added duplicate case test
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        
        if len(actual) != len(c.expected) {
            t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
            continue
        }

        // Check for duplicates
        seen := make(map[string]bool)
        for _, word := range actual {
            if seen[word] {
                t.Errorf("duplicate found: '%v' in input '%v'", word, c.input)
            }
            seen[word] = true
        }

        // Verify words match expected values
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
            }
        }
    }
}
