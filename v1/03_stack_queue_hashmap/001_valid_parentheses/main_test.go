package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "valid ()",
			s:    "()",
			want: true,
		},
		{
			name: "valid ()[]{}",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "invalid (]",
			s:    "(]",
			want: false,
		},
		{
			name: "invalid ([)]",
			s:    "([)]",
			want: false,
		},
		{
			name: "valid {[]}",
			s:    "{[]}",
			want: true,
		},
		{
			name: "empty",
			s:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
