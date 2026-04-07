package main

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		want     int
	}{
		{
			name:    "[2,1,5,6,2,3]",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.heights)
			if got != tt.want {
				t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}
