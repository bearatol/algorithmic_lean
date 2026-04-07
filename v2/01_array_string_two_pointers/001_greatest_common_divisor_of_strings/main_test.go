package main

import (
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name   string
		str1   string
		str2   string
		want   string
	}{
		{
			name:   "ABCABC, ABC",
			str1:   "ABCABC",
			str2:   "ABC",
			want:   "ABC",
		},
		{
			name:   "ABABAB, ABAB",
			str1:   "ABABAB",
			str2:   "ABAB",
			want:   "AB",
		},
		{
			name:   "LEET, CODE",
			str1:   "LEET",
			str2:   "CODE",
			want:   "",
		},
		{
			name:   "ABCDEF, ABC",
			str1:   "ABCDEF",
			str2:   "ABC",
			want:   "",
		},
		{
			name:   "AAAA, AAAA",
			str1:   "AAAA",
			str2:   "AAAA",
			want:   "AAAA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcdOfStrings(tt.str1, tt.str2)
			if got != tt.want {
				t.Errorf("gcdOfStrings(%q, %q) = %q, want %q", tt.str1, tt.str2, got, tt.want)
			}
		})
	}
}
