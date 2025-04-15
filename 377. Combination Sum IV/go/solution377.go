package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}

	var dfs func(t int) int

	dfs = func(t int) int {
		if t == 0 {
			return 1
		}
		if t < 0 {
			return 0
		}
		if dp[t] != -1 {
			return dp[t]
		}
		res := 0
		for _, n := range nums {
			res += dfs(t - n)
		}
		dp[t] = res
		return dp[t]
	}

	return dfs(target)
}
