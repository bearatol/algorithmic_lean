package main

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "hello",
			s:    "hello",
			want: "holle",
		},
		{
			name: "leetcode",
			s:    "leetcode",
			want: "leotcede",
		},
		{
			name: "aA",
			s:    "aA",
			want: "Aa",
		},
		{
			name: "programming",
			s:    "programming",
			want: "prigrammong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowels(tt.s)
			if got != tt.want {
				t.Errorf("reverseVowels(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
