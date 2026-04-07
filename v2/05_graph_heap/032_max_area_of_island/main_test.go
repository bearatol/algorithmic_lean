package main

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name  string
		grid  [][]int
		want  int
	}{
		{
			name:  "[[0,0,1,0,0],[0,0,0,0,0],[0,0,1,0,0],[0,0,0,0,0]]",
			grid:  [][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAreaOfIsland(tt.grid)
			if got != tt.want {
				t.Errorf("maxAreaOfIsland() = %d, want %d", got, tt.want)
			}
		})
	}
}
