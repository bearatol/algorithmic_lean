package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "basic [1,8,6,2,5,4,8,3,7]",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "simple [1,1]",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "decreasing [4,3,2,1,4]",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "two elements [2,3]",
			height: []int{2, 3},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
