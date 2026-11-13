package main

// Append can reuse storage
// Compare append with spare capacity and append beyond capacity. Fill the values observed
// through each alias. Do not assume an exact growth factor.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	spare := make([]int, 1, 3)
	spare[0] = 1
	spareAlias := spare[:3]
	spare = append(spare, 7)

	wantReusedAppended := -999 // TODO: replace this value.
	if got := spareAlias[1]; got != wantReusedAppended {
		t.Errorf("ReusedAppended: got %#v; want %#v", got, wantReusedAppended)
	}

	spare[0] = 9

	wantReusedFirst := -999 // TODO: replace this value.
	if got := spareAlias[0]; got != wantReusedFirst {
		t.Errorf("ReusedFirst: got %#v; want %#v", got, wantReusedFirst)
	}

	full := []int{1}
	fullAlias := full
	grown := append(full, 7)
	grown[0] = 9

	wantAllocatedOld := -999 // TODO: replace this value.
	if got := fullAlias[0]; got != wantAllocatedOld {
		t.Errorf("AllocatedOld: got %#v; want %#v", got, wantAllocatedOld)
	}

	wantAllocatedNew := -999 // TODO: replace this value.
	if got := grown[0]; got != wantAllocatedNew {
		t.Errorf("AllocatedNew: got %#v; want %#v", got, wantAllocatedNew)
	}

	wantAllocatedLength := -999 // TODO: replace this value.
	if got := len(grown); got != wantAllocatedLength {
		t.Errorf("AllocatedLength: got %#v; want %#v", got, wantAllocatedLength)
	}
}
