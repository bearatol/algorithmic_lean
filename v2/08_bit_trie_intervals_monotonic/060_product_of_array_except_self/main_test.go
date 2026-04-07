package main

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "[1,2,3,4]",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("productExceptSelf()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}
