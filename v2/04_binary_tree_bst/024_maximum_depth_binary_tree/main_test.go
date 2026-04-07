package main

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "[3,9,20,null,null,15,7]",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "[1,null,2]",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: 2,
		},
		{
			name: "nil",
			root: nil,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
