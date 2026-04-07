package main

import (
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[1,2,3,1]",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "[2,7,9,3,1]",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
