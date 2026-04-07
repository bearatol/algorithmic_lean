package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "basic case [7,1,5,3,6,4]",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "decreasing prices [7,6,4,3,1]",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single element [1]",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "two elements [1,2]",
			prices: []int{1, 2},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
