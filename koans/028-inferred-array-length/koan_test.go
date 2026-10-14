package main

// Inferred array length
// Complete the array literal using [...] so its length is inferred. Check that its type is
// [3]string and that it contains all three names.

import (
	"fmt"
	"testing"
)

func TestKoan(t *testing.T) {
	names := [...]string{"?"} // TODO: include Ada, Grace, and Edsger.
	if got := fmt.Sprintf("%T", names); got != "[3]string" {
		t.Errorf("type=%s; want [3]string", got)
	}
	if got := fmt.Sprint(names); got != "[Ada Grace Edsger]" {
		t.Errorf("names=%s; want [Ada Grace Edsger]", got)
	}
}
