package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
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
		// if out of boundaries, return 0 (neither profit nor loss)
		if i >= n {
			return 0
		}
		// if calculated before, return it
		if cache[i][bool2int(isBuying)] != -1 {
			return cache[i][bool2int(isBuying)]
		}
		// buy or sell
		res := math.MinInt
		if isBuying {
			// buy current stock
			res = -prices[i] + dfs(i+1, false)
		} else {
			// sell current stock and cooldown next day
			res = prices[i] + dfs(i+2, true)
		}

		// cooldown current stock
		res = max(res, dfs(i+1, isBuying))

		// store result in cache to prevent duplicate calculations
		cache[i][bool2int(isBuying)] = res

		return res
	}

	return dfs(0, true)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
