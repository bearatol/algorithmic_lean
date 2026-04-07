package main

import (
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  int
	}{
		{
			name:  "[1,7,3,6,5,6]",
			nums:  []int{1, 7, 3, 6, 5, 6},
			want:  3,
		},
		{
			name:  "[1,2,3]",
			nums:  []int{1, 2, 3},
			want:  -1,
		},
		{
			name:  "[2,1,-1]",
			nums:  []int{2, 1, -1},
			want:  0,
		},
		{
			name:  "[1]",
			nums:  []int{1},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Errorf("pivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
