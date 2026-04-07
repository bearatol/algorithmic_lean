package main

import (
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "01111011001111",
			s:    "01111011001111",
			want: 8,
		},
		{
			name: "11111",
			s:    "11111",
			want: 4,
		},
		{
			name: "00000",
			s:    "00000",
			want: 0,
		},
		{
			name: "1",
			s:    "1",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarray(tt.s)
			if got != tt.want {
				t.Errorf("longestSubarray(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
