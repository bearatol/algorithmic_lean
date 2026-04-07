package main

import (
	"testing"
)

func TestSmallestInfiniteSet(t *testing.T) {
	s := ConstructorSmallestInfiniteSet()

	// Test sequence: pop 1,2,3, add back 2, pop, add back 1
	tests := []struct {
		name     string
		action   func() int
		expected int
	}{
		{
			name: "first pop",
			action: func() int {
				return s.PopSmallest()
			},
			expected: 1,
		},
		{
			name: "second pop",
			action: func() int {
				return s.PopSmallest()
			},
			expected: 2,
		},
		{
			name: "third pop",
			action: func() int {
				return s.PopSmallest()
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.action()
			if got != tt.expected {
				t.Errorf("action() = %d, want %d", got, tt.expected)
			}
		})
	}
}
