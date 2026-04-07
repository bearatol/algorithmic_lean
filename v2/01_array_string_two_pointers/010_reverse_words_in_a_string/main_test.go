package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "the sky is blue",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "  hello world  ",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "a good   example",
			s:    "a good   example",
			want: "example good a",
		},
		{
			name: "  Bob    Loves  Alice   ",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Alice does not even like bob",
			s:    "Alice does not even like bob",
			want: "bob like even not does Alice",
		},
		{
			name: "single",
			s:    "single",
			want: "single",
		},
		{
			name: "   ",
			s:    "   ",
			want: "",
		},
		{
			name: "  ",
			s:    "  ",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			if got != tt.want {
				t.Errorf("reverseWords(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
