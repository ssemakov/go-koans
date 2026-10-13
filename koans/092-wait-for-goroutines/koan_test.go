package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestKoan(t *testing.T) {
	var wg sync.WaitGroup
	var count atomic.Int64
	for range 3 {
		wg.Add(1)
		go func() { defer wg.Done(); count.Add(1) }()
	}
	wg.Wait()
	want := -999 // TODO: replace this value.
	got := int(count.Load())
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
