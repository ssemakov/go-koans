package main

import (
	"testing"
)

func score(dice []int) int {
	// TODO: implement.
	return -1
}
func TestKoan(t *testing.T) {
	cases := []struct {
		dice []int
		want int
	}{{nil, 0}, {[]int{5}, 50}, {[]int{1}, 100}, {[]int{1, 5, 5, 1}, 300}, {[]int{2, 3, 4, 6}, 0}, {[]int{1, 1, 1}, 1000}, {[]int{2, 2, 2}, 200}, {[]int{3, 3, 3}, 300}, {[]int{4, 4, 4}, 400}, {[]int{5, 5, 5, 5}, 550}, {[]int{6, 6, 6}, 600}, {[]int{1, 1, 1, 5, 1}, 1150}}
	for _, tc := range cases {
		if got := score(tc.dice); got != tc.want {
			t.Fatalf("score(%v) = %d; want %d", tc.dice, got, tc.want)
		}
	}
}
