package main

import (
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		want   string
	}{
		{
			name:   "abc, pqr",
			word1:  "abc",
			word2:  "pqr",
			want:   "apbqcr",
		},
		{
			name:   "ab, pqrs",
			word1:  "ab",
			word2:  "pqrs",
			want:   "apbqrs",
		},
		{
			name:   "abcd, pq",
			word1:  "abcd",
			word2:  "pq",
			want:   "apbqcd",
		},
		{
			name:   "a, b",
			word1:  "a",
			word2:  "b",
			want:   "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("mergeAlternately(%q, %q) = %q, want %q", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}
