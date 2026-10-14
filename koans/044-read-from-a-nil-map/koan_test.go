package main

// Read from a nil map
// Repair putCount so a nil map can receive its first entry. Return the resulting map and
// preserve existing entries when it is already initialized.

import (
	"maps"
	"testing"
)

func putCount(values map[string]int, key string, count int) map[string]int {
	if values == nil {
		return nil
	} // TODO: initialize before inserting.
	values[key] = count
	return values
}

func TestKoan(t *testing.T) {
	var missing map[string]int
	if got := missing["go"]; got != 0 {
		t.Errorf("nil-map read=%d; want 0", got)
	}
	created := putCount(missing, "go", 2)
	if !maps.Equal(created, map[string]int{"go": 2}) {
		t.Errorf("created=%v; want map[go:2]", created)
	}
	existing := map[string]int{"old": 1}
	updated := putCount(existing, "go", 3)
	if !maps.Equal(updated, map[string]int{"old": 1, "go": 3}) {
		t.Errorf("updated=%v", updated)
	}
	if existing["go"] != 3 {
		t.Error("an existing map must be updated in place")
	}
}
