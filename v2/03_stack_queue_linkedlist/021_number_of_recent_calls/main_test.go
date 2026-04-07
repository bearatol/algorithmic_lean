package main

import (
	"testing"
)

func TestRecentCounter(t *testing.T) {
	rc := Constructor()
	tests := []struct {
		name     string
		t        int
		wantMin  int
		wantMax  int
	}{
		{
			name:     "ping 1",
			t:        1,
			wantMin:  1,
			wantMax:  1,
		},
		{
			name:     "ping 100",
			t:        100,
			wantMin:  2,
			wantMax:  2,
		},
		{
			name:     "ping 3001",
			t:        3001,
			wantMin:  3,
			wantMax:  3,
		},
		{
			name:     "ping 3002",
			t:        3002,
			wantMin:  1,
			wantMax:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rc.Ping(tt.t)
			if got < tt.wantMin || got > tt.wantMax {
				t.Errorf("Ping(%d) = %d, want between %d and %d", tt.t, got, tt.wantMin, tt.wantMax)
			}
		})
	}
}
