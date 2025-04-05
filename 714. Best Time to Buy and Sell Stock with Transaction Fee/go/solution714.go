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
	curr := make([]int, 2)
	next := make([]int, 2)

	for i := n - 1; i >= 0; i-- {
		for isBuying := 1; isBuying >= 0; isBuying-- {
			res := math.MinInt
			if isBuying == 1 {
				// buy
				res = -prices[i] + next[0]
			} else {
				// sell
				res = -fee + prices[i] + next[1]
			}

			// cooldown
			res = max(res, next[isBuying])

			curr[isBuying] = res
		}
		next = curr
	}

	return next[1]
}
