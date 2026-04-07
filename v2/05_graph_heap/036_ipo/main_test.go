package main

import (
	"testing"
)

func TestFindMaximizedCapital(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		w        int
		profits  []int
		capital  []int
		want     int
	}{
		{
			name:     "k=2, w=0",
			k:        2,
			w:        0,
			profits:  []int{1, 2, 3},
			capital:  []int{0, 1, 1},
			want:     4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaximizedCapital(tt.k, tt.w, tt.profits, tt.capital)
			if got != tt.want {
				t.Errorf("findMaximizedCapital() = %d, want %d", got, tt.want)
			}
		})
	}
}
