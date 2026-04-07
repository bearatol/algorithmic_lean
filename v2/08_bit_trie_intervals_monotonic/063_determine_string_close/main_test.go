package main

import (
	"testing"
)

func TestHalvesAreAlike(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "book",
			s:    "book",
			want: true,
		},
		{
			name: "textbook",
			s:    "textbook",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := halvesAreAlike(tt.s)
			if got != tt.want {
				t.Errorf("halvesAreAlike(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
