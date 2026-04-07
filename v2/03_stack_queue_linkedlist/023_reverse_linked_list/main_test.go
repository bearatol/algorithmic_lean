package main

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want []int
	}{
		{
			name: "[1,2,3,4,5]",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "[1,2]",
			head: &ListNode{1, &ListNode{2, nil}},
			want: []int{2, 1},
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
			var gotVals []int
			for got != nil {
				gotVals = append(gotVals, got.Val)
				got = got.Next
			}
			if len(gotVals) != len(tt.want) {
				t.Errorf("reverseList() = %v, want %v", gotVals, tt.want)
				return
			}
			for i := range gotVals {
				if gotVals[i] != tt.want[i] {
					t.Errorf("reverseList()[%d] = %d, want %d", i, gotVals[i], tt.want[i])
				}
			}
		})
	}
}
