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
	n0 := 0
	n1 := 0

	for i := n - 1; i >= 0; i-- {
		n1 = max(n1, -prices[i]+n0)
		n0 = max(n0, -fee+prices[i]+n1)
	}

	return n1
}
