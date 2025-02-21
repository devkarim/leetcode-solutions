package main

import "testing"

func TestSolution213(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test Case #1",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "Test Case #2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "Test Case #3",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "Test Case #4",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Test Case #5",
			nums: []int{0, 0},
			want: 0,
		},
		{
			name: "Test Case #6",
			nums: []int{1, 3, 1, 3, 100},
			want: 103,
		},
	}

	for _, tt := range tests {
		actual := rob(tt.nums)

		if actual != tt.want {
			t.Errorf("got %v; expected %v", actual, tt.want)
		}
	}
}
