package main

// Check whether a key exists
// Implement lookupOr. Return the stored value when the key exists, including zero; use
// fallback only for an absent key. Nil maps are allowed.

import (
	"testing"
)

func lookupOr(values map[string]int, key string, fallback int) int {
	return fallback // TODO: distinguish missing from zero.
}

func TestKoan(t *testing.T) {
	counts := map[string]int{"zero": 0, "go": 2}
	cases := []struct {
		values         map[string]int
		key            string
		fallback, want int
	}{
		{counts, "zero", 9, 0}, {counts, "go", 9, 2}, {counts, "missing", 9, 9}, {nil, "go", -1, -1},
	}
	for _, tc := range cases {
		if got := lookupOr(tc.values, tc.key, tc.fallback); got != tc.want {
			t.Errorf("lookupOr(%v,%q,%d)=%d; want %d", tc.values, tc.key, tc.fallback, got, tc.want)
		}
	}
}
