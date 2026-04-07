package main

import (
	"testing"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=4",
			n:    4,
			want: 4,
		},
		{
			name: "n=25",
			n:    25,
			want: 1389537,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tribonacci(tt.n)
			if got != tt.want {
				t.Errorf("tribonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
