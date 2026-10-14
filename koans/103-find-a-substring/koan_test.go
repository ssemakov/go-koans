package main

// Find a substring
// Complete firstNumber using the input text. Find the first run of ASCII digits, parse it,
// and return false when no number is present. Numbers in these inputs fit in int.

import (
	"regexp"
	"strconv"
	"testing"
)

var digits = regexp.MustCompile(`[0-9]+`)

func firstNumber(text string) (int, bool) {
	match := digits.FindString("") // TODO: search the input text.
	value, err := strconv.Atoi(match)
	return value, err == nil
}

func TestKoan(t *testing.T) {
	cases := []struct {
		text string
		want int
		ok   bool
	}{
		{"go123x", 123, true}, {"a7b42", 7, true}, {"zero0", 0, true}, {"x007", 7, true}, {"none", 0, false}, {"", 0, false},
	}
	for _, tc := range cases {
		got, ok := firstNumber(tc.text)
		if got != tc.want || ok != tc.ok {
			t.Errorf("firstNumber(%q)=(%d,%v); want (%d,%v)", tc.text, got, ok, tc.want, tc.ok)
		}
	}
}
