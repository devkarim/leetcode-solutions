package main

import (
	"fmt"
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
		curr[1] = max(curr[1], next[1], -prices[i]+next[0])
		curr[0] = max(curr[0], next[0], -fee+prices[i]+next[1])
		next = curr
	}

	return next[1]
}
