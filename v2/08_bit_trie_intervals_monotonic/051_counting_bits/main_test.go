package main

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n=2",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "n=5",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			if len(got) != len(tt.want) {
				t.Errorf("countBits(%d) = %v, want %v", tt.n, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("countBits()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}
