package main

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "[1,2,3,4,5]",
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "[5,4,3,2,1]",
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := increasingTriplet(tt.nums)
			if got != tt.want {
				t.Errorf("increasingTriplet(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
