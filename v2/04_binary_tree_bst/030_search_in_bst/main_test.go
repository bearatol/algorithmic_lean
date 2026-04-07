package main

import (
	"testing"
)

func TestSearchBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		val  int
		want int
	}{
		{
			name: "found",
			root: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}},
			val:  2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchBST(tt.root, tt.val)
			if got != nil && got.Val != tt.want {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
