package main

// Copy returns the element count
// Fill both the copied counts and complete destinations. copy is limited by the shorter
// length, and it supports overlapping slices.

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	source := []int{1, 2, 3}
	short := []int{9, 9}
	shortCount := copy(short, source)
	long := []int{9, 9, 9, 9}
	longCount := copy(long, source)
	overlap := []int{1, 2, 3, 4}
	overlapCount := copy(overlap[1:], overlap)
	wantCounts := [3]int{-999, -999, -999} // TODO: copied counts.
	wantShort := []int{-999}               // TODO: complete destination.
	wantLong := []int{-999}                // TODO: complete destination.
	wantOverlap := []int{-999}             // TODO: complete destination.
	if got := [3]int{shortCount, longCount, overlapCount}; got != wantCounts {
		t.Errorf("counts=%v; want %v", got, wantCounts)
	}
	if !slices.Equal(short, wantShort) {
		t.Errorf("short=%v; want %v", short, wantShort)
	}
	if !slices.Equal(long, wantLong) {
		t.Errorf("long=%v; want %v", long, wantLong)
	}
	if !slices.Equal(overlap, wantOverlap) {
		t.Errorf("overlap=%v; want %v", overlap, wantOverlap)
	}
}
