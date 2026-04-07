package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "anagram anagram-nagaram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "not anagram rat-car",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "different lengths a-ab",
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
