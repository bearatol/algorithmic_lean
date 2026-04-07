package main

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "basic [1,2,5], amount=11",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "no solution [2], amount=3",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "zero amount",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "[1], amount=0",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
