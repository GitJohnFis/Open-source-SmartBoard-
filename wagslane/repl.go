package main

import (
    "strings"
)

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)

    // Create a map to track seen words
    seen := make(map[string]bool)
    var result []string

    for _, word := range words {
        if !seen[word] { // Add only if it's new
            seen[word] = true
            result = append(result, word)
        }
    }
    return result
}
