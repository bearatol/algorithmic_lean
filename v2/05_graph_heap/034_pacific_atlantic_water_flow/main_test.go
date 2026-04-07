package main

import (
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name     string
		heights  [][]int
		wantRows int
	}{
		{
			name:     "[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,8,4,9,9]]",
			heights:  [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 8, 4, 9, 9}},
			wantRows: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pacificAtlantic(tt.heights)
			if len(got) != tt.wantRows {
				t.Errorf("pacificAtlantic() = %d rows, want %d", len(got), tt.wantRows)
			}
		})
	}
}
