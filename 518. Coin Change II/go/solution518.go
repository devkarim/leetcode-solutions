package main

import "fmt"

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins {
		for total := c; total <= amount; total++ {
			dp[total] += dp[total-c]
		}
	}

	return dp[amount]
}
