package main

import (
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "[1,2,2,1,1,3]",
			arr:  []int{1, 2, 2, 1, 1, 3},
			want: true,
		},
		{
			name: "[1,2,2,3,3,3]",
			arr:  []int{1, 2, 2, 3, 3, 3},
			want: false,
		},
		{
			name: "[-3,0,1,-3,1,1,1,-3,10,0]",
			arr:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueOccurrences(tt.arr)
			if got != tt.want {
				t.Errorf("uniqueOccurrences(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
