package main

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		m    int
		n    int
		want int
	}{
		{
			name: "Grid 3x2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "Grid 3x7",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "Grid 1x1",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "Grid 1x2",
			m:    1,
			n:    2,
			want: 1,
		},
		{
			name: "Grid 2x1",
			m:    2,
			n:    1,
			want: 1,
		},
		{
			name: "Grid 2x2",
			m:    2,
			n:    2,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
