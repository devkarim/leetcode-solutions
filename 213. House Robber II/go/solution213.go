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
		dp := make([]int, end)

		dp[start] = nums[start]
		dp[start+1] = max(dp[start], nums[start+1])

		for i := start + 2; i < end; i++ {
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		}

		return dp[end-1]
	}

	return max(robRange(0, n-1), robRange(1, n))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
