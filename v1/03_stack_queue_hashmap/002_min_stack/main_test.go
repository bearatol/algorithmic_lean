package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	// Создаем стек
	stack := Constructor()

	// Push тесты
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	// GetMin должен вернуть -3
	if got := stack.GetMin(); got != -3 {
		t.Errorf("GetMin() = %d, want -3", got)
	}

	// Pop
	stack.Pop()

	// Top должен вернуть 0
	if got := stack.Top(); got != 0 {
		t.Errorf("Top() = %d, want 0", got)
	}

	// GetMin должен вернуть -2
	if got := stack.GetMin(); got != -2 {
		t.Errorf("GetMin() = %d, want -2", got)
	}
}
