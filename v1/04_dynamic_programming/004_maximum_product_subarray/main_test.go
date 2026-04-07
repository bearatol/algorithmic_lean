package main

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic [2,3,-2,4]",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "[-2,0,-1]",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "[2,3,-2,4,-2]",
			nums: []int{2, 3, -2, 4, -2},
			want: 96,
		},
		{
			name: "single [2]",
			nums: []int{2},
			want: 2,
		},
		{
			name: "all negative [-2,-3,-4]",
			nums: []int{-2, -3, -4},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProduct(tt.nums)
			if got != tt.want {
				t.Errorf("maxProduct(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
