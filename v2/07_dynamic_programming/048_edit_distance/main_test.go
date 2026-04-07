package main

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "horse, ros",
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			name:  "intention, execution",
			word1: "intention",
			word2: "execution",
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("minDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}
