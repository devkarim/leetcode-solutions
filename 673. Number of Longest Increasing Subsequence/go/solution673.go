package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findNumberOfLIS(nums))
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	dp := make([][]int, n) // [[lenLIS, count]]
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	lenLIS, res := 0, 0

	for start := n - 1; start >= 0; start-- {
		maxLen, maxCnt := 1, 1
		for end := start + 1; end < n; end++ {
			if nums[end] <= nums[start] {
				continue
			}
			length, cnt := dp[end][0], dp[end][1]
			if length+1 > maxLen {
				maxLen, maxCnt = length+1, cnt
			} else if length+1 == maxLen {
				maxCnt += cnt
			}
		}
		if maxLen > lenLIS {
			lenLIS, res = maxLen, maxCnt
		} else if maxLen == lenLIS {
			res += maxCnt
		}
		dp[start][0] = maxLen
		dp[start][1] = maxCnt
	}

	return res
}
