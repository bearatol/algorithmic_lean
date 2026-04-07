package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		strs  []string
		want  string
	}{
		{
			name:  "flower,flow,flight",
			strs:  []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  "dog,racecar,car",
			strs:  []string{"dog", "racecar", "car"},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
