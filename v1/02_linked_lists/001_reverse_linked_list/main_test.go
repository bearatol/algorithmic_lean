package main

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "basic [1,2,3,4,5]",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 1},
						},
					},
				},
			},
		},
		{
			name: "single element",
			head: &ListNode{Val: 1},
			want: &ListNode{Val: 1},
		},
		{
			name: "nil",
			head: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.head)
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("reverseList() = %v, want %v", got.Val, tt.want.Val)
					break
				}
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}
