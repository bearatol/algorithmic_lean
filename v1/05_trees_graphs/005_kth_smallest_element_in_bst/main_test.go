package main

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want int
	}{
		{
			name: "basic [5,3,6,2,4,null,null,1], k=3",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:    3,
			want: 3,
		},
		{
			name: "k=1",
			root: &TreeNode{Val: 1},
			k:    1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallest(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("kthSmallest(%d) = %d, want %d", tt.k, got, tt.want)
			}
		})
	}
}
