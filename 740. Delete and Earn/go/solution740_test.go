package main

import "testing"

func Test_deleteAndEarn(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Test Case #1",
			nums: []int{3, 4, 2},
			want: 6,
		},
		{
			name: "Test Case #2",
			nums: []int{2, 2, 3, 3, 3, 4},
			want: 9,
		},
		{
			name: "Test Case #3",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Test Case #4",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "Test Case #5",
			nums: []int{2, 2, 2, 3, 3, 5, 6, 7, 9},
			want: 27,
		},
		{
			name: "Test Case #6",
			nums: []int{1, 1, 1, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteAndEarn(tt.nums)
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
