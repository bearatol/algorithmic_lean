package main

import (
	"testing"
)

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "[3,1,4,3,null,1,5]",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 5},
				},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := goodNodes(tt.root)
			if got != tt.want {
				t.Errorf("goodNodes() = %d, want %d", got, tt.want)
			}
		})
	}
}
