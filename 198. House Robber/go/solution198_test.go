package main

import "testing"

func TestSolution198(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test Case #1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "Test Case #2",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "Test Case #3",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
	}

	for _, tt := range tests {
		actual := rob(tt.nums)

		if actual != tt.want {
			t.Errorf("got %q; expected %q", actual, tt.want)
		}
	}
}
