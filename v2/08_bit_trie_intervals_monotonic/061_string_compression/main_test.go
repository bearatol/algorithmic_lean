package main

import (
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  int
	}{
		{
			name:  "aabcccccaaa",
			chars: []byte("aabcccccaaa"),
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compress(tt.chars)
			if got != tt.want {
				t.Errorf("compress() = %d, want %d", got, tt.want)
			}
		})
	}
}
