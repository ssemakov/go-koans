package main

// Protect shared state with a mutex
// Implement Add and Value using the counter's mutex to protect all accesses. Multiple
// goroutines update and read the same counter. The downloaded test can also be checked
// with go test -race.

import (
	"sync"
	"testing"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Add(delta int) { // TODO: protect and update n.
}
func (c *Counter) Value() int {
	return 0 // TODO: read n under the mutex.
}

func TestKoan(t *testing.T) {
	var counter Counter
	var wg sync.WaitGroup
	for range 200 {
		wg.Add(1)
		go func() { defer wg.Done(); counter.Add(3); counter.Add(-1); _ = counter.Value() }()
	}
	wg.Wait()
	if got := counter.Value(); got != 400 {
		t.Errorf("counter=%d; want 400", got)
	}
	counter.Add(-400)
	if got := counter.Value(); got != 0 {
		t.Errorf("counter after reset=%d; want 0", got)
	}
}
