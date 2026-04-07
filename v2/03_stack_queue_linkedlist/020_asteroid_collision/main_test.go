package main

import (
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name       string
		asteroids  []int
		want       []int
	}{
		{
			name:      "[5,10,-5]",
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			name:      "[8,-8]",
			asteroids: []int{8, -8},
			want:      []int{},
		},
		{
			name:      "[10,2,-5]",
			asteroids: []int{10, 2, -5},
			want:      []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asteroidCollision(tt.asteroids)
			if len(got) != len(tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("asteroidCollision()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}
