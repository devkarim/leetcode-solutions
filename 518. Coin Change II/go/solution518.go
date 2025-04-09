package main

import "fmt"

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	dp[n][0] = 1

	// for total := 0; total <= amount; total++ {
	// 	for i := n - 1; i >= 0; i-- {
	// 		take := 0
	// 		if total >= coins[i] {
	// 			take = dp[i][total-coins[i]]
	// 		}
	// 		skip := dp[i+1][total]
	// 		dp[i][total] = skip + take
	// 	}
	// }

	prev := make([]int, amount+1)
	prev[0] = 1

	for i := n - 1; i >= 0; i-- {
		dp := make([]int, amount+1)
		for total := 0; total <= amount; total++ {
			dp[total] = prev[total]
			if total >= coins[i] {
				dp[total] += dp[total-coins[i]]
			}
		}
		prev = dp
	}

	return prev[amount]
}
