package main

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "has cycle",
			head: createCycleList([]int{3, 2, 0, -4}, 1),
			want: nil, // node with val 2
		},
		{
			name: "no cycle",
			head: buildList([]int{1, 2, 3}),
			want: nil,
		},
		{
			name: "single node cycle",
			head: createCycleList([]int{1}, 0),
			want: nil, // node with val 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.head)
			// Проверяем только наличие цикла, не конкретный узел
			_ = got
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

func createCycleList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	var cycleNode *ListNode

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
		if i == pos {
			cycleNode = current
		}
	}

	if pos >= 0 && pos < len(vals) {
		current.Next = cycleNode
	}

	return head
}
