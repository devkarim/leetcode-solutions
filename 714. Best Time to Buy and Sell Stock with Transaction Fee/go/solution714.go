package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dfs := func(i int, isBuying int) int {
		res := math.MinInt
		if isBuying == 1 {
			// buy
			res = -prices[i] + dp[i+1][0]
		} else {
			// sell
			res = -fee + prices[i] + dp[i+1][1]
		}

		// cooldown
		res = max(res, dp[i+1][isBuying])

		return res
	}

	for i := n - 1; i >= 0; i-- {
		for isBuying := 1; isBuying >= 0; isBuying-- {
			dp[i][isBuying] = dfs(i, isBuying)
		}
	}

	return dp[0][1]
}
