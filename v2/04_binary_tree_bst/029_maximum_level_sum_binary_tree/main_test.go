package main

import (
	"testing"
)

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "[1,7,0,7,-8,null,null]",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: -8}},
				Right: &TreeNode{Val: 0},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxLevelSum(tt.root)
			if got != tt.want {
				t.Errorf("maxLevelSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
