package main

// Clone a map
// Repair cloneGroups. maps.Clone copies entries shallowly; clone each slice too. Preserve
// nil maps and nil slice values, and make edits independent in both directions.

import (
	"maps"
	"slices"
	"testing"
)

func cloneGroups(groups map[string][]int) map[string][]int {
	return maps.Clone(groups) // TODO: also clone the reference-bearing values.
}

func TestKoan(t *testing.T) {
	if got := cloneGroups(nil); got != nil {
		t.Error("nil input should stay nil")
	}
	original := map[string][]int{"go": {1, 2}, "empty": {}, "nil": nil}
	clone := cloneGroups(original)
	if !maps.EqualFunc(original, clone, func(a, b []int) bool { return slices.Equal(a, b) }) {
		t.Fatalf("clone=%v; want %v", clone, original)
	}
	if clone["nil"] != nil {
		t.Error("nil slice became non-nil")
	}
	clone["new"] = []int{7}
	if _, ok := original["new"]; ok {
		t.Error("the cloned map shares entries")
	}
	clone["go"][0] = 9
	if original["go"][0] != 1 {
		t.Error("the cloned slice still shares its array")
	}
	original["go"][1] = 8
	if clone["go"][1] != 2 {
		t.Error("editing the original changed the clone")
	}
}
