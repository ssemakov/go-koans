package main

// Pass values by value
// Fill the caller's values and the returned value. Changing a copied integer or array
// parameter does not update its caller.

import (
	"testing"
)

func increment(value int) int   { value++; return value }
func changeArray(values [2]int) { values[0] = 9 }

func TestKoan(t *testing.T) {
	number := 3
	returned := increment(number)
	array := [2]int{1, 2}
	changeArray(array)

	wantCaller := -999 // TODO: replace this value.
	if got := number; got != wantCaller {
		t.Errorf("Caller: got %#v; want %#v", got, wantCaller)
	}

	wantReturned := -999 // TODO: replace this value.
	if got := returned; got != wantReturned {
		t.Errorf("Returned: got %#v; want %#v", got, wantReturned)
	}

	wantArray := [2]int{-1, -1} // TODO: replace this value.
	if got := array; got != wantArray {
		t.Errorf("Array: got %#v; want %#v", got, wantArray)
	}
}
