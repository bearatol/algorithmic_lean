package main

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want uint32
	}{
		{
			name: "43261596",
			num:  43261596,
			want: 964176192,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBits(tt.num)
			if got != tt.want {
				t.Errorf("reverseBits(%d) = %d, want %d", tt.num, got, tt.want)
			}
		})
	}
}
