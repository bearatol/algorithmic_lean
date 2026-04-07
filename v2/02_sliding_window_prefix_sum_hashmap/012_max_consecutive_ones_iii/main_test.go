package main

import (
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "1101100111, k=2",
			s:    "1101100111",
			k:    2,
			want: 7,
		},
		{
			name: "000000, k=3",
			s:    "000000",
			k:    3,
			want: 3,
		},
		{
			name: "111111, k=1",
			s:    "111111",
			k:    1,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestOnes(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("longestOnes(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
