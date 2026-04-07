package main

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "[10,15,20]",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "[1,100,1,1,1,100,1,1,100,1]",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			if got != tt.want {
				t.Errorf("minCostClimbingStairs(%v) = %d, want %d", tt.cost, got, tt.want)
			}
		})
	}
}
