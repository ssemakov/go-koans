package main

// Wait for goroutines
// Complete parallelMap using one goroutine per value and a WaitGroup. Keep results in
// input order and wait for every worker before returning. The barrier test requires the
// workers to run concurrently.

import (
	"slices"
	"sync"
	"testing"
	"time"
)

func parallelMap(values []int, fn func(int) int) []int {
	result := make([]int, len(values))
	var wg sync.WaitGroup
	_ = wg
	// TODO: launch workers, write each result to its own index, and wait.
	return result
}

func TestKoan(t *testing.T) {
	if got := parallelMap(nil, func(n int) int { return n }); len(got) != 0 {
		t.Errorf("empty result=%v", got)
	}
	values := []int{2, 3, 4}
	started := make(chan int, len(values))
	release := make(chan struct{})
	defer close(release)
	finished := make(chan []int, 1)
	go func() {
		finished <- parallelMap(values, func(n int) int { started <- n; <-release; return n * n })
	}()
	for range values {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("all workers must start before any worker finishes")
		}
	}
	// Signal workers without closing release twice on the failure path.
	for range values {
		release <- struct{}{}
	}
	select {
	case got := <-finished:
		if !slices.Equal(got, []int{4, 9, 16}) {
			t.Errorf("results=%v; want [4 9 16]", got)
		}
	case <-time.After(time.Second):
		t.Fatal("parallelMap did not wait for workers")
	}
	if !slices.Equal(values, []int{2, 3, 4}) {
		t.Error("input changed")
	}
}
