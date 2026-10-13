package main

import (
	"testing"
)

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
func TestKoan(t *testing.T) {
	values := []int{2, 3, 4}
	want := -999 // TODO: replace this value.
	got := sum(values...)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
