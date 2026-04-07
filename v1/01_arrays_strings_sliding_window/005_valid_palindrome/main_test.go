package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "palindrome A man, a plan, a canal: Panama",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "not palindrome race a car",
			s:    "race a car",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single character a",
			s:    "a",
			want: true,
		},
		{
			name: "only alphanumeric 0P",
			s:    "0P",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
