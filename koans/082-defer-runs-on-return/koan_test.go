package main

// Defer runs on return
// Repair runWithCleanup. Register cleanup with defer so it runs once after the work, on
// both success and early return.

import (
	"errors"
	"slices"
	"testing"
)

var ErrWork = errors.New("work failed")

func runWithCleanup(work func() error, cleanup func()) error {
	cleanup() // TODO: run cleanup when this function returns.
	return work()
}

func TestKoan(t *testing.T) {
	for _, fail := range []bool{false, true} {
		var trace []string
		err := runWithCleanup(func() error {
			trace = append(trace, "work")
			if fail {
				return ErrWork
			}
			return nil
		}, func() { trace = append(trace, "cleanup") })
		if !slices.Equal(trace, []string{"work", "cleanup"}) {
			t.Errorf("fail=%v: trace=%q; want [work cleanup]", fail, trace)
		}
		if fail && err != ErrWork {
			t.Errorf("error=%v; want ErrWork", err)
		}
		if !fail && err != nil {
			t.Errorf("error=%v; want nil", err)
		}
	}
}
