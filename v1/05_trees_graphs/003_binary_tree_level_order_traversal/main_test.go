package main

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "basic [3,9,20,null,null,15,7]",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "nil",
			root: nil,
			want: nil,
		},
		{
			name: "single [1]",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("levelOrder() returned %d levels, want %d", len(got), len(tt.want))
			}
		})
	}
}
