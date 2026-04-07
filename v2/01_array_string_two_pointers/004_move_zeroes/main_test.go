package main

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "[0,1,0,3,12]",
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			name:  "[0]",
			input: []int{0},
			want:  []int{0},
		},
		{
			name:  "[0,0,1]",
			input: []int{0, 0, 1},
			want:  []int{1, 0, 0},
		},
		{
			name:  "[1,2,3]",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.input)
			for i := range tt.input {
				if tt.input[i] != tt.want[i] {
					t.Errorf("moveZeroes() = %v, want %v", tt.input, tt.want)
					return
				}
			}
		})
	}
}
