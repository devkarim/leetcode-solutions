package main

import "fmt"

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	n := len(coins)

	dp := make([]int, amount+1)
	dp[0] = 1

	for i := n - 1; i >= 0; i-- {
		for total := 0; total <= amount; total++ {
			if total >= coins[i] {
				dp[total] += dp[total-coins[i]]
			}
		}
	}

	return dp[amount]
}
