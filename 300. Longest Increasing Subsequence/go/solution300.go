package main

import "fmt"

func main() {
	nums := []int{
		10,
		9,
		2,
		5,
		3,
		4,
		101,
		18,
	}

	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	res := 1

	for start := n - 1; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			if nums[end] <= nums[start] {
				continue
			}
			dp[start] = max(dp[start], 1+dp[end])
			res = max(res, dp[start])
		}
	}

	return res
}
