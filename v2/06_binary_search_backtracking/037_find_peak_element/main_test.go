package main

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[1,2,3,1]",
			nums: []int{1, 2, 3, 1},
			want: 2,
		},
		{
			name: "[1,2,1,3,1]",
			nums: []int{1, 2, 1, 3, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.nums)
			if got != tt.want {
				t.Errorf("findPeakElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
