package main

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "23",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "",
			digits: "",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			if len(got) != len(tt.want) {
				t.Errorf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
