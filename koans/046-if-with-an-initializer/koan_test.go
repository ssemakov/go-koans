package main

// If with an initializer
// Complete the if initializer and condition. Use comma-ok so a stored zero succeeds, and
// keep the lookup variable scoped to the if.

import (
	"fmt"
	"testing"
)

func describeLookup(values map[string]int, key string) string {
	if value, ok := values[key]; value != 0 { // TODO: use the presence flag.
		_ = ok
		return fmt.Sprintf("found:%d", value)
	}
	return "missing"
}

func TestKoan(t *testing.T) {
	cases := []struct {
		key  string
		want string
	}{{"zero", "found:0"}, {"go", "found:2"}, {"missing", "missing"}}
	for _, tc := range cases {
		if got := describeLookup(map[string]int{"zero": 0, "go": 2}, tc.key); got != tc.want {
			t.Errorf("key=%q: got %q; want %q", tc.key, got, tc.want)
		}
	}
}
