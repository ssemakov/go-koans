package main

// Declare bit flags
// Complete the first permission constant using 1 << iota. Read, Write, and Execute must
// occupy distinct bits with masks 1, 2, and 4. Keep the other declarations and checks.

import (
	"testing"
)

type Permission uint8

const (
	Read Permission = 0 // TODO: use 1 << iota.
	Write
	Execute
)

func TestKoan(t *testing.T) {
	cases := []struct {
		name       string
		flag, want Permission
	}{
		{"Read", Read, 0b001}, {"Write", Write, 0b010}, {"Execute", Execute, 0b100},
	}
	for _, tc := range cases {
		if tc.flag != tc.want {
			t.Errorf("%s=%08b; want %08b", tc.name, tc.flag, tc.want)
		}
	}
}
