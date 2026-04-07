package main

import (
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int // Количество перестановок
	}{
		{
			name: "[1,2,3]",
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "[0,1]",
			nums: []int{0, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			if len(got) != tt.want {
				t.Errorf("permute() = %d permutations, want %d", len(got), tt.want)
			}
		})
	}
}
