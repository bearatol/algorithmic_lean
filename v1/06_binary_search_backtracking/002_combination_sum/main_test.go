package main

import (
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name      string
		candidates []int
		target    int
		want      [][]int
	}{
		{
			name:      "basic [2,3,6,7], target=7",
			candidates: []int{2, 3, 6, 7},
			target:    7,
			want:      [][]int{{2, 2, 3}, {7}},
		},
		{
			name:      "[2,3,5], target=8",
			candidates: []int{2, 3, 5},
			target:    8,
			want:      [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			// Простая проверка - сравниваем количество комбинаций
			if len(got) != len(tt.want) {
				t.Errorf("combinationSum() returned %d combinations, want %d", len(got), len(tt.want))
			}
			_ = sort.Ints
		})
	}
}
