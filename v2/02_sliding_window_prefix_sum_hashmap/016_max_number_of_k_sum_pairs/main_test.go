package main

import (
	"testing"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		want   int
	}{
		{
			name:   "[1,2,3,4], k=5",
			nums:   []int{1, 2, 3, 4},
			k:      5,
			want:   2,
		},
		{
			name:   "[3,1,3,4,3], k=6",
			nums:   []int{3, 1, 3, 4, 3},
			k:      6,
			want:   1,
		},
		{
			name:   "[2,2,2,2], k=4",
			nums:   []int{2, 2, 2, 2},
			k:      4,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxOperations(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("maxOperations(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
