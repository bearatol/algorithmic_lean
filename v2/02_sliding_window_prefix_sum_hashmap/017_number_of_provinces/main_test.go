package main

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name:        "[[1,1,0],[1,1,0],[0,0,1]]",
			isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			want:        2,
		},
		{
			name:        "[[1,0,0],[0,1,0],[0,0,1]]",
			isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			want:        3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCircleNum(tt.isConnected)
			if got != tt.want {
				t.Errorf("findCircleNum() = %d, want %d", got, tt.want)
			}
		})
	}
}
