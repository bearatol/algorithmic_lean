package main

import (
	"testing"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name  string
		gain  []int
		want  int
	}{
		{
			name:  "[-5,1,5,0,-7]",
			gain:  []int{-5, 1, 5, 0, -7},
			want:  1,
		},
		{
			name:  "[-4,-3,-2,-1,4,3,2]",
			gain:  []int{-4, -3, -2, -1, 4, 3, 2},
			want:  0,
		},
		{
			name:  "[1,2,3]",
			gain:  []int{1, 2, 3},
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestAltitude(tt.gain)
			if got != tt.want {
				t.Errorf("largestAltitude(%v) = %d, want %d", tt.gain, got, tt.want)
			}
		})
	}
}
