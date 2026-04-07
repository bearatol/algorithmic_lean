package main

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		n    int
		want *ListNode
	}{
		{
			name: "basic [1,2,3,4,5], n=2",
			head: buildList([]int{1, 2, 3, 4, 5}),
			n:    2,
			want: buildList([]int{1, 2, 3, 5}),
		},
		{
			name: "remove first [1], n=1",
			head: buildList([]int{1}),
			n:    1,
			want: nil,
		},
		{
			name: "remove second [1,2], n=1",
			head: buildList([]int{1, 2}),
			n:    1,
			want: buildList([]int{1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.head, tt.n)
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("removeNthFromEnd() = %v, want %v", got.Val, tt.want.Val)
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
