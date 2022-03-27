package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.HelloName with a name, checking for a valid
// return value
func TestHelloName(t *testing.T) {
	name := "Alice"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message := HelloName(name)

	if !want.MatchString(message) {
		t.Fatalf(`HelloName(%v) = %q, want match for %#q`, name, message, want)
	}
}
