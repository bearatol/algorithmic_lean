package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "abc, ahbgdc",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "axc, ahbgdc",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "empty, abc",
			s:    "",
			t:    "abc",
			want: true,
		},
		{
			name: "abc, empty",
			s:    "abc",
			t:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isSubsequence(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
