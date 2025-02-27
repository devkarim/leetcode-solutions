package main

import "testing"

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]int
		want int
	}{
		{
			name: "Grid 3x3",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
		{
			name: "Grid 2x4",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 12,
		},
		{
			name: "Grid 1x1",
			grid: [][]int{
				{1},
			},
			want: 1,
		},
		{
			name: "Grid 1x2",
			grid: [][]int{
				{1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPathSum(tt.grid)
			if got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
