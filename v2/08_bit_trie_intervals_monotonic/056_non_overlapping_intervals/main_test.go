package main

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "[[1,2],[2,3],[3,4],[1,3]]",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eraseOverlapIntervals(tt.intervals)
			if got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %d, want %d", got, tt.want)
			}
		})
	}
}
