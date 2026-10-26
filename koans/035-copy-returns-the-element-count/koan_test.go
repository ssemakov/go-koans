package main

// Copy returns the element count
// Predict each copy's return value and the resulting destination. The first two
// destinations are shorter and longer than source. For the overlapping copy, predict the
// entire sharedValues slice.

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	source := []int{1, 2, 3}

	shortDestination := []int{9, 9}
	wantShortCopyCount := -999 // TODO: number of elements copied.
	if copiedCount := copy(shortDestination, source); copiedCount != wantShortCopyCount {
		t.Errorf("copy to shortDestination returned %d; want %d", copiedCount, wantShortCopyCount)
	}
	wantShortDestination := []int{-999} // TODO: complete destination after copy.
	if !slices.Equal(shortDestination, wantShortDestination) {
		t.Errorf("shortDestination=%v; want %v", shortDestination, wantShortDestination)
	}

	longDestination := []int{9, 9, 9, 9}
	wantLongCopyCount := -999 // TODO: number of elements copied.
	if copiedCount := copy(longDestination, source); copiedCount != wantLongCopyCount {
		t.Errorf("copy to longDestination returned %d; want %d", copiedCount, wantLongCopyCount)
	}
	wantLongDestination := []int{-999} // TODO: complete destination after copy.
	if !slices.Equal(longDestination, wantLongDestination) {
		t.Errorf("longDestination=%v; want %v", longDestination, wantLongDestination)
	}

	sharedValues := []int{1, 2, 3, 4}
	overlappingSource := sharedValues[:3]
	overlappingDestination := sharedValues[1:]
	wantOverlapCopyCount := -999 // TODO: number of elements copied.
	if copiedCount := copy(overlappingDestination, overlappingSource); copiedCount != wantOverlapCopyCount {
		t.Errorf("copy to overlappingDestination returned %d; want %d", copiedCount, wantOverlapCopyCount)
	}
	wantSharedValues := []int{-999} // TODO: entire shared slice after copy.
	if !slices.Equal(sharedValues, wantSharedValues) {
		t.Errorf("sharedValues=%v; want %v", sharedValues, wantSharedValues)
	}
}
