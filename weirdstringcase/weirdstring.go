package weirdstringcase

import (
	"strings"
)

// WeirdString converts a string into weird case.
func WeirdString(s string) string {
	parts := strings.Split(s, " ")

	newParts := []string{}

	for _, part := range parts {
		newPart := ""

		for i, rune := range part {
			if i%2 == 0 {
				newPart += strings.ToUpper(string(rune))
			} else {
				newPart += strings.ToLower(string(rune))
			}
		}

		newParts = append(newParts, newPart)
	}

	return strings.Join(newParts, " ")
}
