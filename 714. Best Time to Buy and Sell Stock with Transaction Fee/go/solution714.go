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
	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, 2)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i int, isBuying bool) int

	dfs = func(i int, isBuying bool) int {
		if i >= n {
			return 0
		}
		if cache[i][bool2int(isBuying)] != -1 {
			return cache[i][bool2int(isBuying)]
		}

		res := math.MinInt
		if isBuying {
			// buy
			res = -prices[i] + dfs(i+1, false)
		} else {
			// sell
			res = -fee + prices[i] + dfs(i+1, true)
		}

		// cooldown
		res = max(res, dfs(i+1, isBuying))

		cache[i][bool2int(isBuying)] = res

		return cache[i][bool2int(isBuying)]
	}

	return dfs(0, true)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
