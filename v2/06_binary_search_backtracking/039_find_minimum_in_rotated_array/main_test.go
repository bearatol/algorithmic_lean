package main

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[3,4,5,1,2]",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "[4,5,6,7,0,1,2]",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.want {
				t.Errorf("findMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
