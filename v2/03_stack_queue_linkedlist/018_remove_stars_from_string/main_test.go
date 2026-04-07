package main

import (
	"testing"
)

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "erase*",
			s:    "erase*****",
			want: "",
		},
		{
			name: "abc*",
			s:    "abc*d*c",
			want: "ab",
		},
		{
			name: "abcd",
			s:    "abcd",
			want: "abcd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeStars(tt.s)
			if got != tt.want {
				t.Errorf("removeStars(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
