package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	robRange := func(start, end int) int {
		beforeBefore := 0
		before := nums[start]

		for i := start + 1; i < end; i++ {
			beforeBefore, before = before, max(nums[i]+beforeBefore, before)
		}

		return before
	}

	return max(robRange(0, n-1), robRange(1, n))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
