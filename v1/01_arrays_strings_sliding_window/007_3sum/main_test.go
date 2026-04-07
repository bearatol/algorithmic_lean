package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "basic [-1,0,1,2,-1,-4]",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "no solution [1,2,-2,-1]",
			nums: []int{1, 2, -2, -1},
			want: nil,
		},
		{
			name: "empty",
			nums: []int{},
			want: nil,
		},
		{
			name: "single element [0]",
			nums: []int{0},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			// Сортируем для сравнения
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) != len(got[j]) {
					return len(got[i]) < len(got[j])
				}
				for k := range got[i] {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
