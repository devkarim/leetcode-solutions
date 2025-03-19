package main

import "fmt"

func main() {
	nums := []int{9, 4, 7, 2, 10}
	fmt.Println(longestArithSeqLength(nums))
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)

	dp := make([]map[int]int, n)
	res := 2

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[j] - nums[i]
			if val, ok := dp[j][diff]; ok {
				dp[i][diff] = val + 1
			} else {
				dp[i][diff] = 2
			}
			res = max(res, dp[i][diff])
		}
	}

	return res
}
