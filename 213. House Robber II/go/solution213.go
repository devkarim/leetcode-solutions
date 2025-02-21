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
	if n == 2 {
		return max(nums[0], nums[1])
	}

	robRange := func(start, end int) int {
		beforeBefore := nums[start]
		before := max(beforeBefore, nums[start+1])

		for i := start + 2; i < end; i++ {
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
