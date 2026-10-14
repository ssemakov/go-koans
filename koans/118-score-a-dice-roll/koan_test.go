package main

// Score a dice roll
// Implement score for zero to five dice, each from 1 through 6. Three ones score 1000;
// other triples score 100 times the face. Remaining ones score 100, fives score 50, and
// other faces score zero. Order must not matter; preserve the input.

import (
	"slices"
	"testing"
)

func score(dice []int) int {
	return -1 // TODO: count faces and score each die once.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		dice []int
		want int
	}{
		{nil, 0}, {[]int{}, 0}, {[]int{5}, 50}, {[]int{1}, 100}, {[]int{1, 5, 5, 1}, 300},
		{[]int{2, 3, 4, 6}, 0}, {[]int{1, 1, 1}, 1000}, {[]int{1, 1, 1, 1}, 1100}, {[]int{1, 1, 1, 1, 1}, 1200},
		{[]int{2, 2, 2}, 200}, {[]int{3, 3, 3}, 300}, {[]int{4, 4, 4}, 400},
		{[]int{5, 5, 5}, 500}, {[]int{5, 5, 5, 5}, 550}, {[]int{5, 5, 5, 5, 5}, 600}, {[]int{6, 6, 6}, 600},
		{[]int{1, 1, 1, 5, 1}, 1150}, {[]int{3, 5, 3, 1, 3}, 450}, {[]int{2, 5, 2, 1, 2}, 350},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.dice)
		if got := score(tc.dice); got != tc.want {
			t.Errorf("score(%v)=%d; want %d", before, got, tc.want)
		}
		if !slices.Equal(tc.dice, before) {
			t.Error("score changed its input")
		}
		reversed := slices.Clone(before)
		slices.Reverse(reversed)
		if got := score(reversed); got != tc.want {
			t.Errorf("score reversed(%v)=%d; want %d", before, got, tc.want)
		}
	}
}
