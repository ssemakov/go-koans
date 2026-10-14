package main

// If and else
// Implement sign with if/else. Return negative, zero, or positive according to the input.

import (
	"testing"
)

func sign(value int) string {
	return "" // TODO: implement with if/else.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		value int
		want  string
	}{{-7, "negative"}, {-1, "negative"}, {0, "zero"}, {1, "positive"}, {8, "positive"}}
	for _, tc := range cases {
		if got := sign(tc.value); got != tc.want {
			t.Errorf("sign(%d)=%q; want %q", tc.value, got, tc.want)
		}
	}
}
