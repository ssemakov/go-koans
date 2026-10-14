package main

// Parse a time
// Repair parseDate's layout using Go's reference date. Parse YYYY-MM-DD in UTC, and
// preserve errors for invalid dates or a different input format.

import (
	"testing"
	"time"
)

func parseDate(text string) (time.Time, error) {
	return time.Parse("YYYY-MM-DD", text) // TODO: use the Go reference-date layout.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		text    string
		year    int
		month   time.Month
		day     int
		invalid bool
	}{
		{"2024-03-14", 2024, time.March, 14, false}, {"2000-02-29", 2000, time.February, 29, false},
		{"2023-02-29", 0, 0, 0, true}, {"14-03-2024", 0, 0, 0, true}, {"2024-3-14", 0, 0, 0, true}, {"", 0, 0, 0, true},
	}
	for _, tc := range cases {
		value, err := parseDate(tc.text)
		if tc.invalid {
			if err == nil {
				t.Errorf("parseDate(%q) should fail", tc.text)
			}
			continue
		}
		if err != nil {
			t.Errorf("parseDate(%q): %v", tc.text, err)
			continue
		}
		year, month, day := value.Date()
		if year != tc.year || month != tc.month || day != tc.day || value.Location() != time.UTC {
			t.Errorf("date=%v; want %04d-%02d-%02d UTC", value, tc.year, tc.month, tc.day)
		}
	}
}
