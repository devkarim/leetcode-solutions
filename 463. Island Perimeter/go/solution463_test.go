package main

import "testing"

func Test_islandPerimeter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]int
		want int
	}{
		{
			name: "Test Case #1",
			grid: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			want: 16,
		},
		{
			name: "Test Case #2",
			grid: [][]int{
				{1},
			},
			want: 4,
		},
		{
			name: "Test Case #3",
			grid: [][]int{
				{1, 0},
			},
			want: 4,
		},
		{
			name: "Test Case #4",
			grid: [][]int{
				{0, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := islandPerimeter(tt.grid)
			if got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
