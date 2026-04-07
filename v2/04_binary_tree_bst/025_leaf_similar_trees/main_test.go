package main

import (
	"testing"
)

func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		name  string
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		{
			name: "same leaves",
			root1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			root2: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leafSimilar(tt.root1, tt.root2)
			if got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
