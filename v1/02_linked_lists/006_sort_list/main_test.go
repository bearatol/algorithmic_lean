package main

import (
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "basic [4,2,1,3]",
			head: buildList([]int{4, 2, 1, 3}),
			want: buildList([]int{1, 2, 3, 4}),
		},
		{
			name: "already sorted [-1,5,3,4,0]",
			head: buildList([]int{-1, 5, 3, 4, 0}),
			want: buildList([]int{-1, 0, 3, 4, 5}),
		},
		{
			name: "single element",
			head: buildList([]int{1}),
			want: buildList([]int{1}),
		},
		{
			name: "nil",
			head: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortList(tt.head)
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("sortList() = %v, want %v", got.Val, tt.want.Val)
					break
				}
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}
