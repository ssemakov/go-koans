package main

// Named types
// Implement addBonus. Score is a defined type, so convert between Score and int explicitly
// when combining them.

import (
	"testing"
)

type Score int

func addBonus(score Score, bonus int) Score {
	return score // TODO: combine the values with an explicit conversion.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		score Score
		bonus int
		want  Score
	}{{4, 3, 7}, {0, 5, 5}, {10, -2, 8}}
	for _, tc := range cases {
		if got := addBonus(tc.score, tc.bonus); got != tc.want {
			t.Errorf("addBonus(%d,%d)=%d; want %d", tc.score, tc.bonus, got, tc.want)
		}
	}
}
