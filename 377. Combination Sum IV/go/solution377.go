package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for t := 1; t <= target; t++ {
		for _, n := range nums {
			if t >= n {
				dp[t] += dp[t-n]
			}
		}
	}

	return dp[target]
}
