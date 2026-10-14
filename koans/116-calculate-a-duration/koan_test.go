package main

// Calculate a duration
// Implement billableMinutes for intervals up to 24 hours. Round each positive partial
// minute up, keep zero at zero, and return ErrOrder when end is before start. Compare
// instants even if their timezones differ.

import (
	"errors"
	"testing"
	"time"
)

var ErrOrder = errors.New("end is before start")

func billableMinutes(start, end time.Time) (int, error) {
	return 0, nil // TODO: calculate and round elapsed minutes.
}

func TestKoan(t *testing.T) {
	start := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		end  time.Time
		want int
		err  error
	}{
		{start, 0, nil}, {start.Add(time.Second), 1, nil}, {start.Add(time.Minute), 1, nil},
		{start.Add(61 * time.Second), 2, nil}, {start.Add(90*time.Minute + 30*time.Second), 91, nil},
		{start.Add(5 * time.Minute).In(time.FixedZone("offset", 2*60*60)), 5, nil},
		{start.Add(-time.Second), 0, ErrOrder},
	}
	for _, tc := range cases {
		minutes, err := billableMinutes(start, tc.end)
		if minutes != tc.want || !errors.Is(err, tc.err) {
			t.Errorf("end=%v: got (%d,%v); want (%d,%v)", tc.end, minutes, err, tc.want, tc.err)
		}
	}
}
