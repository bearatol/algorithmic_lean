package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "3[a]2[bc]",
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			name: "3[a2[c]]",
			s:    "3[a2[c]]",
			want: "accaccacc",
		},
		{
			name: "2[abc]3[cd]ef",
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeString(tt.s)
			if got != tt.want {
				t.Errorf("decodeString(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
