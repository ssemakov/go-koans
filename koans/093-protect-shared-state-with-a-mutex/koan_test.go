package main

import (
	"sync"
	"testing"
)

func TestKoan(t *testing.T) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for range 5 {
		wg.Add(1)
		go func() { defer wg.Done(); mu.Lock(); count++; mu.Unlock() }()
	}
	wg.Wait()
	want := -999 // TODO: replace this value.
	got := count
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
