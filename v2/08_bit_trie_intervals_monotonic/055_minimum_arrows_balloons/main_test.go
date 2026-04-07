package main

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "[[10,16],[2,8],[1,6],[7,12]]",
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinArrowShots(tt.points)
			if got != tt.want {
				t.Errorf("findMinArrowShots() = %d, want %d", got, tt.want)
			}
		})
	}
}
