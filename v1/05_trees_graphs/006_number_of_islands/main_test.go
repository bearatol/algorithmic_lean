package main

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name  string
		grid  [][]byte
		want  int
	}{
		{
			name: "basic",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "multiple islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "no islands",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.grid)
			if got != tt.want {
				t.Errorf("numIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}
