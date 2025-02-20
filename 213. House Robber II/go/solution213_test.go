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
	}

	for _, tt := range tests {
		actual := rob(tt.nums)

		if actual != tt.want {
			t.Errorf("got %v; expected %v", actual, tt.want)
		}
	}
}
