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

	dfs := func(i, total int) int {
		if i == n {
			if total == 0 {
				return 1
			}
			return 0
		}
		take := 0
		if total >= coins[i] {
			take = dp[i][total-coins[i]]
		}
		skip := dp[i+1][total]
		return skip + take
	}

	for total := 0; total <= amount; total++ {
		for i := n; i >= 0; i-- {
			dp[i][total] = dfs(i, total)
		}
	}

	return dp[0][amount]
}
