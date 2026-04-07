package main

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		name       string
		root       *TreeNode
		targetSum  int
		want       int
	}{
		{
			name:       "[10,5,-3,3,2,null,11,3,-2,null,1], targetSum=8",
			root:       nil, // Создайте дерево
			targetSum:  8,
			want:       3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathSum(tt.root, tt.targetSum)
			if got != tt.want {
				t.Errorf("pathSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
