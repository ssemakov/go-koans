package main

// Boolean zero value
// Complete the Enabled initialization. Leave the omitted flags at their zero values; set
// Enabled from the input rather than a constant.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	for _, requested := range []bool{false, true} {
		config := struct{ Enabled, Verbose bool }{
			Enabled: false, // TODO: use the requested setting.
		}
		if config.Enabled != requested {
			t.Errorf("requested=%v: Enabled=%v", requested, config.Enabled)
		}
		if config.Verbose {
			t.Error("an omitted boolean field should be false")
		}
	}
}
