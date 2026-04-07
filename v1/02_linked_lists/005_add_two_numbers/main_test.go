package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "342 + 465 = 807",
			l1:   buildList([]int{2, 4, 3}),
			l2:   buildList([]int{5, 6, 4}),
			want: buildList([]int{7, 0, 8}),
		},
		{
			name: "0 + 0 = 0",
			l1:   buildList([]int{0}),
			l2:   buildList([]int{0}),
			want: buildList([]int{0}),
		},
		{
			name: "9999999 + 1 = 10000000",
			l1:   buildList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   buildList([]int{1}),
			want: buildList([]int{0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got.Val, tt.want.Val)
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
