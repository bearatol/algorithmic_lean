package main

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "3x7",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "3x2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "1x1",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "1x2",
			m:    1,
			n:    2,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
