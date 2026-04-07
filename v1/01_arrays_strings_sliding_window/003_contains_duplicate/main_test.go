package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "has duplicate [1,2,3,1]",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "no duplicate [1,2,3,4]",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "all same [1,1,1,1]",
			nums: []int{1, 1, 1, 1},
			want: true,
		},
		{
			name: "empty []",
			nums: []int{},
			want: false,
		},
		{
			name: "single [1]",
			nums: []int{1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
