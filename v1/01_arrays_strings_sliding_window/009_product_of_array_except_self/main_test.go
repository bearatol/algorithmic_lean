package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic [1,2,3,4]",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "with zero [1,2,0,3,4]",
			nums: []int{1, 2, 0, 3, 4},
			want: []int{0, 0, 24, 0, 0},
		},
		{
			name: "two zeros [0,1,0,3,4]",
			nums: []int{0, 1, 0, 3, 4},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "single [1]",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "two [1,2]",
			nums: []int{1, 2},
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
