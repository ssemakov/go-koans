package main

// Multiple return values
// Complete the call assignments to first. Receive its value and presence flag
// together; a stored zero is present. In the final call, discard the value with _
// and keep only the flag. Keep first and the checks unchanged.

import "testing"

func first(values []int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}
	return values[0], true
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		want   int
		ok     bool
	}{
		{nil, 0, false}, {[]int{}, 0, false}, {[]int{0, 7}, 0, true}, {[]int{4, 5}, 4, true},
	}
	for _, tc := range cases {
		got, ok := 0, false // TODO: receive both results from first(tc.values).
		if got != tc.want || ok != tc.ok {
			t.Errorf("first(%v)=(%d,%v); want (%d,%v)", tc.values, got, ok, tc.want, tc.ok)
		}
	}
	present := false // TODO: call first([]int{0}), discarding its first result.
	if !present {
		t.Error("a stored zero is present")
	}
}
