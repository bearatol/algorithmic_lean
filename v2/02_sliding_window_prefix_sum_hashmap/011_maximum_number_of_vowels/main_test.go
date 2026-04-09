package main

import (
	"testing"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "rhythms, k=4",
			s:    "rhythms",
			k:    4,
			want: 0,
		},
		{
			name: "abciiidef, k=3",
			s:    "abciiidef",
			k:    3,
			want: 3,
		},
		{
			name: "aeiou, k=2",
			s:    "aeiou",
			k:    2,
			want: 2,
		},
		{
			name: "tryhard, k=4",
			s:    "tryhard",
			k:    4,
			want: 1,
		},
		{
			name: "qwertyuiop, k=5",
			s:    "qwertyuiop",
			k:    5,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxVowels(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("maxVowels(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
