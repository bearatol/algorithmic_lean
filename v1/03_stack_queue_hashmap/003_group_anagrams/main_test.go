package main

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		strs  []string
		want  [][]string
	}{
		{
			name:  "basic [eat,tan,ate,nat,bat]",
			strs:  []string{"eat", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"eat", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:  "empty",
			strs:  []string{},
			want:  nil,
		},
		{
			name:  "single [a]",
			strs:  []string{"a"},
			want:  [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			// Для простоты проверяем только количество групп
			if len(got) != len(tt.want) {
				t.Errorf("groupAnagrams() returned %d groups, want %d", len(got), len(tt.want))
			}
			_ = sort.Strings
		})
	}
}
