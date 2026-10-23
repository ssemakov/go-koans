package main

// Append and shared storage
// Fill the complete expected slices after both appends. Keep the setup and append
// operations unchanged. Compare capacities 4 and 2: can first change without being
// reassigned?

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	cases := []struct {
		capacity   int
		wantBase   []int
		wantFirst  []int
		wantSecond []int
	}{
		{4, []int{-999}, []int{-999}, []int{-999}}, // TODO: base, first, second.
		{2, []int{-999}, []int{-999}, []int{-999}}, // TODO: base, first, second.
	}
	for _, tc := range cases {
		base := make([]int, 2, tc.capacity)
		base[0] = 10
		base[1] = 20

		first := append(base, 30)
		second := append(base, 40)

		if !slices.Equal(base, tc.wantBase) {
			t.Errorf("capacity=%d: base=%v; want %v", tc.capacity, base, tc.wantBase)
		}
		if !slices.Equal(first, tc.wantFirst) {
			t.Errorf("capacity=%d: first=%v; want %v", tc.capacity, first, tc.wantFirst)
		}
		if !slices.Equal(second, tc.wantSecond) {
			t.Errorf("capacity=%d: second=%v; want %v", tc.capacity, second, tc.wantSecond)
		}
	}
}
