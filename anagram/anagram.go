/*
Package anagram exposes functions to work with anagrams.
*/
package anagram

import (
	"sort"
	"strings"
)

// arrange sorts the letters in a string ascendingly and returns the sorted
// string.
func arrange(s string) string {
	letters := strings.Split(s, "")

	sort.Strings(letters)

	return strings.Join(letters, "")
}

// normalize returns a trimmed and lowercased string.
func normalize(s string) string {
	return strings.ToLower(strings.Trim(s, " "))
}

// CheckAnagramInSlice checks for possible anagrams of a given word in a string
// slice.
func CheckAnagramInSlice(w string, words []string) []string {
	result := []string{}

	normalizedW := normalize(w)
	comparableW := arrange(normalizedW)

	for _, word := range words {
		normalizedWord := normalize(word)
		comparableWord := arrange(normalizedWord)

		if comparableW == comparableWord && normalizedW != normalizedWord {
			result = append(result, word)
		}
	}

	return result
}
