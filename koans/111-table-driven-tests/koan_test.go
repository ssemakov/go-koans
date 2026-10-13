package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	cases := []struct{ a, b, sum int }{{1, 2, 3}, {2, 3, 5}}
	total := 0
	for _, tc := range cases {
		if tc.a+tc.b == tc.sum {
			total++
		}
	}
	want := -999 // TODO: replace this value.
	got := total
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
