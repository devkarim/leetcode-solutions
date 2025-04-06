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

	buy := math.MinInt
	sell := 0
	prevSell := 0
	prevBuy := 0

	for i := n - 1; i >= 0; i-- {
		prevBuy = buy
		buy = max(-prices[i]+prevSell, prevBuy)

		prevSell = sell
		sell = max(prices[i]+prevBuy, prevSell)
	}

	return sell
}
