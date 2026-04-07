package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "abcabcbb",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "bbbbb",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "pwwkew",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty",
			s:    "",
			want: 0,
		},
		{
			name: "dvdf",
			s:    "dvdf",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
