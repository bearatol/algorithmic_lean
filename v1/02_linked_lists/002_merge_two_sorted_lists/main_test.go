package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		list1  *ListNode
		list2  *ListNode
		want   *ListNode
	}{
		{
			name:   "basic [1,2,4] + [1,3,4]",
			list1:  buildList([]int{1, 2, 4}),
			list2:  buildList([]int{1, 3, 4}),
			want:   buildList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:   "both nil",
			list1:  nil,
			list2:  nil,
			want:   nil,
		},
		{
			name:   "one nil",
			list1:  buildList([]int{0}),
			list2:  nil,
			want:   buildList([]int{0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("mergeTwoLists() = %v, want %v", got.Val, tt.want.Val)
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
