package main

import "testing"

// AI generated test cases
func TestCleanInputRocket(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {input: "  ", expected: []string{}},                // Empty input
        {input: "  hello  ", expected: []string{"hello"}},  // Single word with spaces
        {input: "  hello hello world  ", expected: []string{"hello", "world"}},  // Check duplicate removal
        {input: "  HellO  World  ", expected: []string{"hello", "world"}},  // Case normalization test
    }

    for _, c := range cases {
        actual := cleanInput(c.input) // Ensure cleanInput is correctly implemented

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

        // Validate word order and correctness
        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
            }
        }
    }
}
