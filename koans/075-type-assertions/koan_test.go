package main

// Type assertions
// Implement stringValue with a comma-ok type assertion. Return the string and true only
// for dynamic type string, including an empty string. Other types must return empty and
// false without panicking.

import (
	"testing"
)

type Label string

func stringValue(value any) (string, bool) {
	return "", false // TODO: use a safe type assertion.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		value any
		want  string
		ok    bool
	}{
		{"Go", "Go", true}, {"", "", true}, {7, "", false}, {nil, "", false}, {Label("Go"), "", false}, {[]byte("Go"), "", false},
	}
	for _, tc := range cases {
		got, ok := stringValue(tc.value)
		if got != tc.want || ok != tc.ok {
			t.Errorf("stringValue(%T)=(%q,%v); want (%q,%v)", tc.value, got, ok, tc.want, tc.ok)
		}
	}
}
