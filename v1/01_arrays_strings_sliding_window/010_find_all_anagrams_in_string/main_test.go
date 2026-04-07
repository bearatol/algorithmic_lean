package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "cbaebabacbabc",
			s:    "cbaebabacbabc",
			p:    "abc",
			want: []int{0, 6, 12},
		},
		{
			name: "abab",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
		{
			name: "no anagrams",
			s:    "hello",
			p:    "world",
			want: nil,
		},
		{
			name: "empty p",
			s:    "abc",
			p:    "",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
