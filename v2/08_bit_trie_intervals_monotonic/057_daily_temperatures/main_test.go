package main

import (
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "[73,74,75,71,69,72,76,73]",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if len(got) != len(tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("dailyTemperatures()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}
