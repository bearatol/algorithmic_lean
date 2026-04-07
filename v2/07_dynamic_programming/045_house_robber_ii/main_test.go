package main

import (
	"testing"
)

func TestRob2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[2,3,2]",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "[1,2,3,1]",
			nums: []int{1, 2, 3, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob2(tt.nums)
			if got != tt.want {
				t.Errorf("rob2(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
