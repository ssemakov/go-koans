package main

// Switch without an expression
// Repair the expressionless switch. Test invalid values first, then high scores. Return
// invalid outside 0..100, pass for 60..100, and fail otherwise.

import (
	"testing"
)

func grade(score int) string {
	switch {
	case score >= 60:
		return "pass" // TODO: put invalid inputs first.
	case score < 0 || score > 100:
		return "invalid"
	default:
		return "fail"
	}
}

func TestKoan(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{{-1, "invalid"}, {0, "fail"}, {59, "fail"}, {60, "pass"}, {100, "pass"}, {101, "invalid"}}
	for _, tc := range cases {
		if got := grade(tc.score); got != tc.want {
			t.Errorf("grade(%d)=%q; want %q", tc.score, got, tc.want)
		}
	}
}
