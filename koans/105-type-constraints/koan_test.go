package main

import (
	"testing"
)

func sum[T ~int | ~float64](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}
func TestKoan(t *testing.T) {

	want := -999 // TODO: replace this value.
	got := sum([]int{2, 3})
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
