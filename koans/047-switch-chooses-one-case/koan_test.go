package main

// Switch chooses one case
// Implement commandName with switch. Start and run share a case; stop has another; any
// other input returns unknown. Go cases do not fall through automatically.

import (
	"testing"
)

func commandName(command string) string {
	return "" // TODO: implement a switch.
}

func TestKoan(t *testing.T) {
	cases := []struct{ command, want string }{{"start", "running"}, {"run", "running"}, {"stop", "stopped"}, {"", "unknown"}, {"restart", "unknown"}}
	for _, tc := range cases {
		if got := commandName(tc.command); got != tc.want {
			t.Errorf("commandName(%q)=%q; want %q", tc.command, got, tc.want)
		}
	}
}
