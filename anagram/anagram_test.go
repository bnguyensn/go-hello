package anagram

import (
	"reflect"
	"testing"
)

func TestCheckAnagramInSlice(t *testing.T) {
	s1 := "hicam"
	slice1 := []string{
		"Andrew",
		"Micha",
		"junior",
	}

	expected := []string{
		"Micha",
	}
	received := CheckAnagramInSlice(s1, slice1)

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf(`CheckAnagramInSlice(%v, %v) = %v, expected %v`, s1, slice1, received, expected)
	}
}
