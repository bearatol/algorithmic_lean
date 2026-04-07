package main

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name    string
		text1   string
		text2   string
		want    int
	}{
		{
			name:  "abcde, ace",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			name:  "abc, def",
			text1: "abc",
			text2: "def",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.want {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}
