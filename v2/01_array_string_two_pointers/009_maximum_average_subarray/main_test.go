package main

import (
	"math"
	"testing"
)

func TestMaximumAverageSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{
			name: "[1,12,-5,-6,50,3], k=4",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		},
		{
			name: "[5], k=1",
			nums: []int{5},
			k:    1,
			want: 5.0,
		},
		{
			name: "[0,4,0,3,2], k=1",
			nums: []int{0, 4, 0, 3, 2},
			k:    1,
			want: 4.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumAverageSubarray(tt.nums, tt.k)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("maximumAverageSubarray(%v, %d) = %f, want %f", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
