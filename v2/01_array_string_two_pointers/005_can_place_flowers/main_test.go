package main

import (
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{
			name:      "[1,0,0,0,1], n=1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "[1,0,0,0,1], n=2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		{
			name:      "[0,0,1,0,1], n=1",
			flowerbed: []int{0, 0, 1, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "[0], n=1",
			flowerbed: []int{0},
			n:         1,
			want:      true,
		},
		{
			name:      "[1,0,0,0,0,1], n=2",
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPlaceFlowers(tt.flowerbed, tt.n)
			if got != tt.want {
				t.Errorf("canPlaceFlowers(%v, %d) = %v, want %v", tt.flowerbed, tt.n, got, tt.want)
			}
		})
	}
}
