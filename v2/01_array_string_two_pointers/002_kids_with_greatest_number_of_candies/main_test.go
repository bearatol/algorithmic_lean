package main

import (
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name        string
		candies     []int
		extraCandies int
		want        []bool
	}{
		{
			name:        "[2,3,5,1,3], extraCandies=3",
			candies:     []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:        []bool{true, true, true, false, true},
		},
		{
			name:        "[4,2,1,1,2], extraCandies=1",
			candies:     []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:        []bool{true, false, false, false, false},
		},
		{
			name:        "[12,1,12], extraCandies=10",
			candies:     []int{12, 1, 12},
			extraCandies: 10,
			want:        []bool{true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kidsWithCandies(tt.candies, tt.extraCandies)
			if len(got) != len(tt.want) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("kidsWithCandies()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
