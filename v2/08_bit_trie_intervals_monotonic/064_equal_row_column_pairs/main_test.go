package main

import (
	"testing"
)

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "[[3,2,1],[1,7,6],[2,7,7]]",
			grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equalPairs(tt.grid)
			if got != tt.want {
				t.Errorf("equalPairs() = %d, want %d", got, tt.want)
			}
		})
	}
}
